import { PlainMessage } from '@bufbuild/protobuf'
import { AiFillGithub, AiFillGitlab } from 'solid-icons/ai'
import { SiGitea } from 'solid-icons/si'
import { JSXElement } from 'solid-js'
import {
  Application,
  Application_ContainerState,
  BuildStatus,
  CreateWebsiteRequest,
  DeployType,
  PortPublicationProtocol,
  Website,
} from '/@/api/neoshowcase/protobuf/gateway_pb'
import { colorVars } from '/@/theme'

export const buildStatusStr: Record<BuildStatus, string> = {
  [BuildStatus.QUEUED]: 'Queued',
  [BuildStatus.BUILDING]: 'Building',
  [BuildStatus.SUCCEEDED]: 'Succeeded',
  [BuildStatus.FAILED]: 'Failed',
  [BuildStatus.CANCELLED]: 'Cancelled',
  [BuildStatus.SKIPPED]: 'Skipped',
}

export enum ApplicationState {
  Idle = 'Idle',
  Deploying = 'Deploying',
  Running = 'Running',
  Serving = 'Serving',
  Error = 'Error',
}

export const deploymentState = (app: Application): ApplicationState => {
  if (!app.running) {
    return ApplicationState.Idle
  }
  if (app.currentBuild === '') {
    // First build may still be running
    return ApplicationState.Idle
  }
  if (app.deployType === DeployType.RUNTIME) {
    switch (app.container) {
      case Application_ContainerState.MISSING:
      case Application_ContainerState.STARTING:
        return ApplicationState.Deploying
      case Application_ContainerState.RUNNING:
        return ApplicationState.Running
      case Application_ContainerState.RESTARTING:
      case Application_ContainerState.EXITED:
      case Application_ContainerState.ERRORED:
      case Application_ContainerState.UNKNOWN:
        return ApplicationState.Error
    }
  }
  return ApplicationState.Serving
}

const errorCommit = '0'.repeat(40)

export const applicationState = (app: Application): ApplicationState => {
  if (!app.running) {
    return ApplicationState.Idle
  }
  if (app.commit === errorCommit) {
    return ApplicationState.Error
  }
  switch (app.latestBuildStatus) {
    case BuildStatus.QUEUED:
      return ApplicationState.Deploying
    case BuildStatus.BUILDING:
      return ApplicationState.Deploying
    case BuildStatus.SUCCEEDED:
      return deploymentState(app)
    case BuildStatus.FAILED:
      return ApplicationState.Error
    case BuildStatus.CANCELLED:
      return deploymentState(app)
    case BuildStatus.SKIPPED:
      return deploymentState(app)
    case undefined:
      return ApplicationState.Error
  }
}

export type Provider = 'GitHub' | 'GitLab' | 'Gitea'

export const repositoryURLToProvider = (url: string): Provider => {
  const normalizedURL = url.toLowerCase()
  if (normalizedURL.includes('github')) return 'GitHub'
  if (normalizedURL.includes('gitlab')) return 'GitLab'
  if (normalizedURL.includes('gitea')) return 'Gitea'
  if (normalizedURL.includes('git.trap.jp')) return 'Gitea'
  return 'GitHub' // fallback?
}

export const providerToIcon = (provider: Provider, size = 20): JSXElement => {
  switch (provider) {
    case 'GitHub':
      return <AiFillGithub size={size} color={colorVars.semantic.text.black} />
    case 'GitLab':
      return <AiFillGitlab size={size} color="#FC6D26" />
    case 'Gitea':
      return <SiGitea size={size} color={colorVars.semantic.text.black} />
  }
}

export const getWebsiteURL = (website: PlainMessage<Website | CreateWebsiteRequest>): string => {
  const scheme = website.https ? 'https' : 'http'
  return `${scheme}://${website.fqdn}${website.pathPrefix}`
}

export const websiteWarnings = (subdomain: string | undefined): string[] => {
  const warnings = []
  if (subdomain?.includes('_')) {
    warnings.push('アンダースコア「_」を含むホスト名は非推奨です。将来非対応になる可能性があります。')
  }
  const labels = subdomain?.split('.')
  if (labels && labels.length >= 2) {
    warnings.push('このホスト名では専用の証明書が取得されます。可能な限りホストのラベル数は少なくしてください。')
  }
  return warnings
}

export const extractRepositoryNameFromURL = (url: string): string => {
  const segments = url.split('/')
  const lastSegment = segments.pop() || segments.pop() // 末尾のスラッシュを除去
  return lastSegment?.replace(/\.git$/, '') ?? ''
}

export const portPublicationProtocolMap: Record<PortPublicationProtocol, string> = {
  [PortPublicationProtocol.TCP]: 'TCP',
  [PortPublicationProtocol.UDP]: 'UDP',
}
