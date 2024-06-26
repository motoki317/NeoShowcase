import { type Component, Match, Show, Switch } from 'solid-js'
import type {
  ApplicationConfig,
  BuildConfigRuntimeBuildpack,
  BuildConfigRuntimeCmd,
  BuildConfigRuntimeDockerfile,
  BuildConfigStaticBuildpack,
  BuildConfigStaticCmd,
  BuildConfigStaticDockerfile,
  RuntimeConfig,
  StaticConfig,
} from '/@/api/neoshowcase/protobuf/gateway_pb'
import Code from '/@/components/UI/Code'
import { List } from '../List'

const BuildpackConfigInfo: Component<{ config: BuildConfigRuntimeBuildpack | BuildConfigStaticBuildpack }> = (
  props,
) => {
  return (
    <List.Row>
      <List.RowContent>
        <List.RowTitle>Context</List.RowTitle>
        <Code value={props.config.context} />
      </List.RowContent>
    </List.Row>
  )
}
const CmdConfigInfo: Component<{ config: BuildConfigRuntimeCmd | BuildConfigStaticCmd }> = (props) => {
  return (
    <>
      <List.Row>
        <List.RowContent>
          <List.RowTitle>Base Image</List.RowTitle>
          <List.RowData>{props.config.baseImage}</List.RowData>
        </List.RowContent>
      </List.Row>
      <Show when={props.config.buildCmd !== ''}>
        <List.Row>
          <List.RowContent>
            <List.RowTitle>Build Command</List.RowTitle>
            <Code value={props.config.buildCmd} />
          </List.RowContent>
        </List.Row>
      </Show>
    </>
  )
}
const DockerfileConfigInfo: Component<{ config: BuildConfigRuntimeDockerfile | BuildConfigStaticDockerfile }> = (
  props,
) => {
  return (
    <>
      <List.Row>
        <List.RowContent>
          <List.RowTitle>Dockerfile</List.RowTitle>
          <List.RowData>{props.config.dockerfileName}</List.RowData>
        </List.RowContent>
      </List.Row>
      <List.Row>
        <List.RowContent>
          <List.RowTitle>Context</List.RowTitle>
          <Code value={props.config.context} />
        </List.RowContent>
      </List.Row>
    </>
  )
}

const RuntimeConfigInfo: Component<{ config?: RuntimeConfig }> = (props) => {
  return (
    <Show when={props.config}>
      <List.Columns>
        <List.Row>
          <List.RowContent>
            <List.RowTitle>Use MariaDB</List.RowTitle>
            <List.RowData>{`${props.config!.useMariadb}`}</List.RowData>
          </List.RowContent>
        </List.Row>
        <List.Row>
          <List.RowContent>
            <List.RowTitle>Use MongoDB</List.RowTitle>
            <List.RowData>{`${props.config!.useMongodb}`}</List.RowData>
          </List.RowContent>
        </List.Row>
      </List.Columns>
      <Show when={props.config!.entrypoint !== ''}>
        <List.Row>
          <List.RowContent>
            <List.RowTitle>Entrypoint</List.RowTitle>
            <Code value={props.config!.entrypoint} />
          </List.RowContent>
        </List.Row>
      </Show>
      <Show when={props.config!.command !== ''}>
        <List.Row>
          <List.RowContent>
            <List.RowTitle>Command</List.RowTitle>
            <Code value={props.config!.command} />
          </List.RowContent>
        </List.Row>
      </Show>
    </Show>
  )
}
const StaticConfigInfo: Component<{ config?: StaticConfig }> = (props) => {
  return (
    <Show when={props.config}>
      <List.Row>
        <List.RowContent>
          <List.RowTitle>Artifact Path</List.RowTitle>
          <Code value={props.config!.artifactPath} />
        </List.RowContent>
      </List.Row>
      <List.Row>
        <List.RowContent>
          <List.RowTitle>Single Page Application</List.RowTitle>
          <List.RowData>{`${props.config!.spa}`}</List.RowData>
        </List.RowContent>
      </List.Row>
    </Show>
  )
}

const AppConfigInfo: Component<{ config: ApplicationConfig }> = (props) => {
  const c = props.config.buildConfig
  return (
    <Switch>
      <Match when={c.case === 'runtimeBuildpack' && c}>
        {(c) => (
          <>
            <BuildpackConfigInfo config={c().value} />
            <RuntimeConfigInfo config={c().value.runtimeConfig} />
          </>
        )}
      </Match>
      <Match when={c.case === 'runtimeCmd' && c}>
        {(c) => (
          <>
            <CmdConfigInfo config={c().value} />
            <RuntimeConfigInfo config={c().value.runtimeConfig} />
          </>
        )}
      </Match>
      <Match when={c.case === 'runtimeDockerfile' && c}>
        {(c) => (
          <>
            <DockerfileConfigInfo config={c().value} />
            <RuntimeConfigInfo config={c().value.runtimeConfig} />
          </>
        )}
      </Match>
      <Match when={c.case === 'staticBuildpack' && c}>
        {(c) => (
          <>
            <BuildpackConfigInfo config={c().value} />
            <StaticConfigInfo config={c().value.staticConfig} />
          </>
        )}
      </Match>
      <Match when={c.case === 'staticCmd' && c}>
        {(c) => (
          <>
            <CmdConfigInfo config={c().value} />
            <StaticConfigInfo config={c().value.staticConfig} />
          </>
        )}
      </Match>
      <Match when={c.case === 'staticDockerfile' && c}>
        {(c) => (
          <>
            <DockerfileConfigInfo config={c().value} />
            <StaticConfigInfo config={c().value.staticConfig} />
          </>
        )}
      </Match>
    </Switch>
  )
}

export default AppConfigInfo
