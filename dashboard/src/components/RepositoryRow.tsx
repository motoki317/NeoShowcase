import { JSXElement } from 'solid-js'
import {
  application,
  appsCount,
  container,
  header,
  headerLeft,
  addBranchButton,
  repoName,
  appName,
  appDetail,
  appFooter,
  appFooterRight,
  applicationNotLast,
} from '/@/components/Repository.css'
import { AiFillGithub, AiFillGitlab } from 'solid-icons/ai'
import { vars } from '/@/theme.css'
import { SiGitea } from 'solid-icons/si'
import { StatusIcon } from '/@/components/StatusIcon'
import { Application, Repository } from '/@/api/neoshowcase/protobuf/apiserver_pb'
import { applicationState, repositoryURLToProvider } from '/@/libs/application'
import { durationHuman, shortSha } from '/@/libs/format'

export type Provider = 'GitHub' | 'GitLab' | 'Gitea'

export interface Props {
  repo: Repository
  apps: Application[]
}

const providerToIcon = (provider: Provider): JSXElement => {
  switch (provider) {
    case 'GitHub':
      return <AiFillGithub size={20} color={vars.text.black1} />
    case 'GitLab':
      return <AiFillGitlab size={20} color='#FC6D26' />
    case 'Gitea':
      return <SiGitea size={20} color={vars.text.black1} />
  }
}

export const RepositoryRow = ({ repo, apps }: Props): JSXElement => {
  const provider = repositoryURLToProvider(repo.url)
  return (
    <div class={container}>
      <div class={header}>
        <div class={headerLeft}>
          {providerToIcon(provider)}
          <div class={repoName}>{repo.name}</div>
          <div class={appsCount}>
            {apps.length} {apps.length === 1 ? 'app' : 'apps'}
          </div>
        </div>
        <div class={addBranchButton}>
          <div>Add&nbsp;branch</div>
        </div>
      </div>
      {apps.map((app, i) => (
        <div class={i === apps.length - 1 ? application : applicationNotLast}>
          <StatusIcon state={applicationState(app)} />
          <div class={appDetail}>
            <div class={appName}>{app.name}</div>
            <div class={appFooter}>
              <div>{shortSha(app.currentCommit)}</div>
              <div class={appFooterRight}>
                <div>{app.websites[0]?.fqdn || ''}</div>
                <div>{durationHuman(3 * 60 * 1000) /* TODO: use updatedAt */}</div>
              </div>
            </div>
          </div>
        </div>
      ))}
    </div>
  )
}
