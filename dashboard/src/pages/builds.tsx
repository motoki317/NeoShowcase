import { Title } from '@solidjs/meta'
import { type Component, Show, createMemo, createResource } from 'solid-js'
import {
  type Application,
  GetApplicationsRequest_Scope,
  GetRepositoriesRequest_Scope,
  Repository,
} from '../api/neoshowcase/protobuf/gateway_pb'
import { MaterialSymbols } from '../components/UI/MaterialSymbols'
import { MainViewContainer } from '../components/layouts/MainView'
import { WithNav } from '../components/layouts/WithNav'
import { BuildList, List } from '../components/templates/List'
import { Nav } from '../components/templates/Nav'
import { client, getRepositoryCommits } from '../libs/api'

const builds: Component = () => {
  const [apps] = createResource(() => client.getApplications({ scope: GetApplicationsRequest_Scope.ALL }))

  const appMap = (): Record<string, Application> => {
    const a = apps()
    if (!a) return {}
    return Object.fromEntries(a.applications.map((a) => [a.id, a]))
  }

  const [builds] = createResource(() =>
    client
      .getAllBuilds({
        limit: 100,
      })
      .then((res) => res.builds),
  )
  const hashes = () => builds()?.map((b) => b.commit)
  const [commits] = createResource(() => getRepositoryCommits(hashes() || []))

  const sortedBuilds = createMemo(
    () =>
      builds()
        ?.sort((b1, b2) => {
          return (b2.queuedAt?.toDate().getTime() ?? 0) - (b1.queuedAt?.toDate().getTime() ?? 0)
        })
        ?.map((b) => ({ build: b, app: appMap()[b.applicationId] })) ?? [],
  )
  const showPlaceHolder = () => builds()?.length === 0

  return (
    <WithNav.Container>
      <Title>Build Queue - NeoShowcase</Title>
      <WithNav.Navs>
        <Nav title="Build Queue" />
      </WithNav.Navs>
      <WithNav.Body>
        <MainViewContainer background="grey">
          <Show when={showPlaceHolder()} fallback={<BuildList builds={sortedBuilds()} commits={commits()} />}>
            <List.Container>
              <List.PlaceHolder>
                <MaterialSymbols displaySize={80}>deployed_code</MaterialSymbols>
                No Builds
              </List.PlaceHolder>
            </List.Container>
          </Show>
        </MainViewContainer>
      </WithNav.Body>
    </WithNav.Container>
  )
}

export default builds
