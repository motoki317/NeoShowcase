import { Timestamp } from '@bufbuild/protobuf'
import { Code, ConnectError } from '@connectrpc/connect'
import { styled } from '@macaron-css/solid'
import { Component, For, Show, createEffect, createMemo, createResource, createSignal, onCleanup } from 'solid-js'
import { ApplicationOutput } from '/@/api/neoshowcase/protobuf/gateway_pb'
import { Button } from '/@/components/UI/Button'
import { LogContainer } from '/@/components/UI/LogContainer'
import { client, handleAPIError } from '/@/libs/api'
import { toWithAnsi } from '/@/libs/buffers'
import { isScrolledToBottom } from '/@/libs/scroll'
import { addTimestamp, lessTimestamp, minTimestamp } from '/@/libs/timestamp'

const LoadMoreContainer = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    gap: '8px',
    marginBottom: '6px',
    fontSize: '16px',
  },
})

const loadLimitSeconds = 7 * 86400
const loadDuration = 86400n

const loadLogChunk = async (appID: string, before: Timestamp, limit: number): Promise<ApplicationOutput[]> => {
  const res = await client.getOutput({ applicationId: appID, before: before, limit: limit })
  return res.outputs
}

const oldestTimestamp = (ts: ApplicationOutput[]): Timestamp =>
  ts.reduce((acc, t) => (t.time ? minTimestamp(acc, t.time) : acc), Timestamp.now())
const sortByTimestamp = (ts: ApplicationOutput[]) =>
  ts.sort((a, b) => (a.time && b.time ? (lessTimestamp(a.time, b.time) ? -1 : 1) : 0))

export interface ContainerLogProps {
  appID: string
  showTimestamp: boolean
}

export const ContainerLog: Component<ContainerLogProps> = (props) => {
  const [loadedUntil, setLoadedUntil] = createSignal(Timestamp.now())
  const [olderLogs, setOlderLogs] = createSignal<ApplicationOutput[]>([])

  const loadDisabled = () => Timestamp.now().seconds - loadedUntil().seconds >= loadLimitSeconds
  const [loading, setLoading] = createSignal(false)
  const load = async () => {
    setLoading(true)
    try {
      const loadedOlderLogs = await loadLogChunk(props.appID, loadedUntil(), 100)
      if (loadedOlderLogs.length === 0) {
        setLoadedUntil(addTimestamp(loadedUntil(), -loadDuration))
      } else {
        setLoadedUntil(oldestTimestamp(loadedOlderLogs))
      }
      sortByTimestamp(loadedOlderLogs)
      setOlderLogs((prev) => loadedOlderLogs.concat(prev))
    } catch (e) {
      handleAPIError(e, 'ログの読み込み中にエラーが発生しました')
    }
    setLoading(false)
  }

  const logStreamAbort = new AbortController()
  const [logStream] = createResource(
    () => props.appID,
    (id) => client.getOutputStream({ id }, { signal: logStreamAbort.signal }),
  )
  const [streamedLog, setStreamedLog] = createSignal<ApplicationOutput[]>([])
  createEffect(() => {
    const stream = logStream()
    if (!stream) {
      setStreamedLog([])
      return
    }

    const iterate = async () => {
      try {
        for await (const log of stream) {
          setStreamedLog((prev) => prev.concat(log))
        }
      } catch (err) {
        // ignore abort error
        const isAbortErr = err instanceof ConnectError && err.code === Code.Canceled
        if (!isAbortErr) {
          console.trace(err)
          return
        }
      }
    }

    void iterate()
  })
  onCleanup(() => {
    logStreamAbort.abort()
  })

  const streamedLogOldest = createMemo(() => {
    const logs = streamedLog()
    if (logs.length === 0) return
    return logs.reduce((acc, log) => (log.time ? minTimestamp(acc, log.time) : acc), Timestamp.now())
  })
  createEffect(() => {
    const oldest = streamedLogOldest()
    if (!oldest) return
    if (lessTimestamp(oldest, loadedUntil())) {
      setLoadedUntil(oldest)
    }
  })

  let logRef: HTMLDivElement
  createEffect(() => {
    streamedLog()
    const ref = logRef
    if (!ref) return
    if (atBottom()) {
      ref.scrollTop = ref.scrollHeight
    }
  })

  const [atBottom, setAtBottom] = createSignal(true)
  const onScroll = (e: { target: Element }) => setAtBottom(isScrolledToBottom(e.target))

  return (
    <LogContainer ref={logRef!} overflowX="scroll" onScroll={onScroll}>
      {/* cannot distinguish zero log and loading (but should be enough for most use-cases) */}
      <Show when={streamedLog().length > 0}>
        <LoadMoreContainer>
          Loaded until {loadedUntil().toDate().toLocaleString()}
          <Show when={!loadDisabled()} fallback={<span>(reached load limit)</span>}>
            <Button variants="ghost" size="small" onClick={load} disabled={loading()}>
              {loading() ? 'Loading...' : 'Load more'}
            </Button>
          </Show>
        </LoadMoreContainer>
      </Show>
      <For each={olderLogs()}>{(log) => <code innerHTML={formatLogLine(log, props.showTimestamp)} />}</For>
      <For each={streamedLog()}>{(log) => <code innerHTML={formatLogLine(log, props.showTimestamp)} />}</For>
    </LogContainer>
  )
}

const formatLogLine = (log: ApplicationOutput, withTimestamp: boolean): string => {
  return (withTimestamp ? `${log.time?.toDate().toLocaleString()} ` : '') + toWithAnsi(log.log)
}
