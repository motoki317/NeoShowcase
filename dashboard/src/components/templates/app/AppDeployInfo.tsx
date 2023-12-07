import { styled } from '@macaron-css/solid'
import { Component, For, Show, createSignal } from 'solid-js'
import toast from 'solid-toast'
import { Application, Build, DeployType, Repository } from '/@/api/neoshowcase/protobuf/gateway_pb'
import Badge from '/@/components/UI/Badge'
import { Button } from '/@/components/UI/Button'
import JumpButton from '/@/components/UI/JumpButton'
import { ToolTip } from '/@/components/UI/ToolTip'
import { URLText } from '/@/components/UI/URLText'
import { client, handleAPIError } from '/@/libs/api'
import { ApplicationState, deploymentState, getWebsiteURL } from '/@/libs/application'
import { titleCase } from '/@/libs/casing'
import { colorOverlay } from '/@/libs/colorOverlay'
import { diffHuman, shortSha } from '/@/libs/format'
import { colorVars, media, textVars } from '/@/theme'
import { List } from '../List'
import { AppStatusIcon } from './AppStatusIcon'

const DeploymentContainer = styled('div', {
  base: {
    width: '100%',
    display: 'grid',
    gridTemplateColumns: 'repeat(3, 1fr)',
    gridTemplateRows: 'auto',
    gap: '1px',

    background: colorVars.semantic.ui.border,
    border: `1px solid ${colorVars.semantic.ui.border}`,
    borderRadius: '8px',
    overflow: 'hidden',

    '@media': {
      [media.mobile]: {
        gridTemplateColumns: 'repeat(2, 1fr)',
      },
    },
  },
})
const AppStateContainer = styled('div', {
  base: {
    position: 'relative',
    gridArea: '1 / 1 / 4 / 2',
    width: '100%',
    display: 'grid',
    gridTemplateRows: '1fr 2fr 1fr',
    justifyItems: 'center',

    cursor: 'pointer',
    color: colorVars.semantic.text.black,
    ...textVars.h3.medium,

    '@media': {
      [media.mobile]: {
        gridArea: '1 / 1 / 2 / 3',
      },
    },
  },
  variants: {
    variant: {
      Running: {
        background: colorOverlay(colorVars.semantic.ui.primary, colorVars.semantic.transparent.successSelected),
        selectors: {
          '&:hover': {
            background: colorOverlay(colorVars.semantic.ui.primary, colorVars.semantic.transparent.successHover),
          },
        },
      },
      Static: {
        background: colorOverlay(colorVars.semantic.ui.primary, colorVars.semantic.transparent.primarySelected),
        selectors: {
          '&:hover': {
            background: colorOverlay(colorVars.semantic.ui.primary, colorVars.semantic.transparent.primaryHover),
          },
        },
      },
      Idle: {
        background: colorOverlay(colorVars.semantic.ui.primary, colorVars.primitive.blackAlpha[200]),
        selectors: {
          '&:hover': {
            background: colorOverlay(colorVars.semantic.ui.primary, colorVars.primitive.blackAlpha[100]),
          },
        },
      },
      Deploying: {
        background: colorOverlay(colorVars.semantic.ui.primary, colorVars.semantic.transparent.warnSelected),
        selectors: {
          '&:hover': {
            background: colorOverlay(colorVars.semantic.ui.primary, colorVars.semantic.transparent.warnHover),
          },
        },
      },
      Error: {
        background: colorOverlay(colorVars.semantic.ui.primary, colorVars.semantic.transparent.errorSelected),
        selectors: {
          '&:hover': {
            background: colorOverlay(colorVars.semantic.ui.primary, colorVars.semantic.transparent.errorHover),
          },
        },
      },
    },
  },
})
const AppState = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    gap: '8px',
  },
})
const ActionButtons = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    gap: '8px',
  },
})
const DeployInfoContainer = styled('div', {
  base: {
    width: '100%',
    padding: '16px 20px',
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    gap: '8px',

    background: colorVars.semantic.ui.primary,
  },
  variants: {
    long: {
      true: {
        gridColumn: 'span 2',
      },
    },
  },
})
const AppDeployInfo: Component<{
  app: Application
  refetchApp: () => void
  repo: Repository
  refreshRepo: () => void
  disableRefresh: () => boolean
  deployedBuild: Build | undefined
  latestBuildId: string | undefined
  hasPermission: boolean
}> = (props) => {
  const [mouseEnter, setMouseEnter] = createSignal(false)
  const showActions = () => props.hasPermission && mouseEnter()

  const restartApp = async () => {
    try {
      await client.startApplication({ id: props.app.id })
      await props.refetchApp()
      toast.success('アプリケーションを再起動しました')
    } catch (e) {
      handleAPIError(e, 'アプリケーションの再起動に失敗しました')
    }
  }
  const stopApp = async () => {
    try {
      await client.stopApplication({ id: props.app.id })
      await props.refetchApp()
      toast.success('アプリケーションを停止しました')
    } catch (e) {
      handleAPIError(e, 'アプリケーションの停止に失敗しました')
    }
  }

  return (
    <DeploymentContainer>
      <AppStateContainer
        onMouseEnter={() => setMouseEnter(true)}
        onMouseLeave={() => setMouseEnter(false)}
        variant={deploymentState(props.app)}
      >
        <div />
        <AppState>
          <AppStatusIcon state={deploymentState(props.app)} size={80} />
          {deploymentState(props.app)}
        </AppState>
        <Show when={showActions()}>
          <ActionButtons>
            <Button variants="borderError" size="small" onClick={restartApp} disabled={props.disableRefresh()}>
              {props.app.running ? 'Restart App' : 'Start App'}
            </Button>
            <Button
              variants="borderError"
              size="small"
              onClick={stopApp}
              disabled={props.disableRefresh() || deploymentState(props.app) === ApplicationState.Idle}
            >
              Stop App
            </Button>
          </ActionButtons>
        </Show>
      </AppStateContainer>
      <DeployInfoContainer>
        <List.RowContent>
          <List.RowTitle>起動時刻</List.RowTitle>
          <Show when={props.app.updatedAt}>
            {(nonNullUpdatedAt) => {
              const { diff, localeString } = diffHuman(nonNullUpdatedAt().toDate())

              return (
                <ToolTip props={{ content: localeString }}>
                  <List.RowData>{diff}</List.RowData>
                </ToolTip>
              )
            }}
          </Show>
        </List.RowContent>
      </DeployInfoContainer>
      <DeployInfoContainer>
        <List.RowContent>
          <List.RowTitle>Deploy Type</List.RowTitle>
          <List.RowData>{titleCase(DeployType[props.app.deployType])}</List.RowData>
        </List.RowContent>
        <JumpButton href={`/apps/${props.app.id}/settings/build`} />
      </DeployInfoContainer>
      <DeployInfoContainer long>
        <List.RowContent>
          <List.RowTitle>Source Commit</List.RowTitle>
          <List.RowData>
            {`${props.deployedBuild?.commit ? shortSha(props.deployedBuild?.commit) : '0000000'}`}
            <Show when={props.deployedBuild?.id === props.latestBuildId}>
              <ToolTip props={{ content: '最新のビルドがデプロイされています' }}>
                <Badge variant="success">Latest</Badge>
              </ToolTip>
            </Show>
          </List.RowData>
        </List.RowContent>
        <Show when={props.hasPermission}>
          <Button
            variants="ghost"
            size="medium"
            onClick={props.refreshRepo}
            disabled={props.disableRefresh()}
            tooltip={{
              props: {
                content: 'リポジトリの最新コミットを取得',
              },
            }}
          >
            Refresh Commit
          </Button>
        </Show>
      </DeployInfoContainer>
      <DeployInfoContainer long>
        <List.RowContent>
          <List.RowTitle>
            Domains
            <Badge variant="text">{props.app.websites.length}</Badge>
          </List.RowTitle>
          <For each={props.app.websites.map(getWebsiteURL)}>
            {(url) => (
              <List.RowData>
                <URLText text={url} href={url} />
              </List.RowData>
            )}
          </For>
        </List.RowContent>
        <JumpButton href={`/apps/${props.app.id}/settings/domains`} />
      </DeployInfoContainer>
    </DeploymentContainer>
  )
}

export default AppDeployInfo
