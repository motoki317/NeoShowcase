import { styled } from '@macaron-css/solid'
import { Outlet, useMatch, useNavigate } from '@solidjs/router'
import { ErrorBoundary, Show, Suspense, useTransition } from 'solid-js'
import { Button } from '/@/components/UI/Button'
import { MaterialSymbols } from '/@/components/UI/MaterialSymbols'
import ErrorView from '/@/components/layouts/ErrorView'
import { MainViewContainer } from '/@/components/layouts/MainView'
import { SideView } from '/@/components/layouts/SideView'
import SuspenseContainer from '/@/components/layouts/SuspenseContainer'
import SettingSkeleton from '/@/components/templates/SettingSkeleton'
import { useRepositoryData } from '/@/routes'

const SideMenu = styled('div', {
  base: {
    position: 'sticky',
    width: '100%',
    top: '0',
    display: 'flex',
    flexDirection: 'column',
  },
})

export default () => {
  const { repo } = useRepositoryData()
  const loaded = () => !!repo()
  const matchGeneralPage = useMatch(() => `/repos/${repo()?.id}/settings/`)
  const matchAuthPage = useMatch(() => `/repos/${repo()?.id}/settings/authorization`)
  const matchOwnerPage = useMatch(() => `/repos/${repo()?.id}/settings/owner`)

  const [isPending, start] = useTransition()
  const navigator = useNavigate()
  const navigate = (path: string) => start(() => navigator(path))

  return (
    <MainViewContainer>
      <Show when={loaded()}>
        <SideView.Container>
          <SideView.Side>
            <SideMenu>
              <Button
                variants="text"
                size="medium"
                full
                active={!!matchGeneralPage()}
                onclick={() => {
                  navigate(`/repos/${repo()?.id}/settings/`)
                }}
                leftIcon={<MaterialSymbols>browse_activity</MaterialSymbols>}
              >
                General
              </Button>
              <Button
                variants="text"
                size="medium"
                full
                active={!!matchAuthPage()}
                onclick={() => {
                  navigate(`/repos/${repo()?.id}/settings/authorization`)
                }}
                leftIcon={<MaterialSymbols>conversion_path</MaterialSymbols>}
              >
                Authorization
              </Button>
              <Button
                variants="text"
                size="medium"
                full
                active={!!matchOwnerPage()}
                onclick={() => {
                  navigate(`/repos/${repo()?.id}/settings/owner`)
                }}
                leftIcon={<MaterialSymbols>person</MaterialSymbols>}
              >
                Owner
              </Button>
            </SideMenu>
          </SideView.Side>
          <SideView.Main>
            <ErrorBoundary fallback={(props) => <ErrorView {...props} />}>
              <Suspense fallback={<SettingSkeleton />}>
                <SuspenseContainer isPending={isPending()}>
                  <Outlet />
                </SuspenseContainer>
              </Suspense>
            </ErrorBoundary>
          </SideView.Main>
        </SideView.Container>
      </Show>
    </MainViewContainer>
  )
}
