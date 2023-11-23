import { createPromiseClient } from '@connectrpc/connect'
import { createConnectTransport } from '@connectrpc/connect-web'
import { createResource } from 'solid-js'
import toast from 'solid-toast'
import { APIService } from '/@/api/neoshowcase/protobuf/gateway_connect'

const transport = createConnectTransport({
  baseUrl: '',
})
export const client = createPromiseClient(APIService, transport)

export const [user] = createResource(() => client.getMe({}))
export const [systemInfo] = createResource(() => client.getSystemInfo({}))
export const [availableMetrics] = createResource(() => client.getAvailableMetrics({}))

export const handleAPIError = (e, message: string) => {
  if (e.message) {
    //' e instanceof ConnectError' does not work for some reason
    toast.error(`${message}\n${e.message}`)
  } else {
    console.trace(e)
    toast.error('予期しないエラーが発生しました')
  }
}
