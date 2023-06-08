import { useParams } from '@solidjs/router'
import { Component, JSX, Show, createEffect, createMemo, createResource, createSignal, onMount } from 'solid-js'
import { client, handleAPIError } from '/@/libs/api'
import { Container } from '/@/libs/layout'
import { Header } from '/@/components/Header'
import { AppNav } from '/@/components/AppNav'
import { Button } from '/@/components/Button'
import { styled } from '@macaron-css/solid'
import { vars } from '/@/theme'
import { createStore } from 'solid-js/store'
import {
  ApplicationConfig,
  BuildConfigRuntimeBuildpack,
  BuildConfigRuntimeCmd,
  BuildConfigRuntimeDockerfile,
  BuildConfigStaticCmd,
  BuildConfigStaticDockerfile,
  CreateWebsiteRequest,
  PortPublication,
  RuntimeConfig,
  UpdateApplicationRequest,
  User,
} from '/@/api/neoshowcase/protobuf/gateway_pb'
import { BuildConfigs } from '/@/components/BuildConfigs'
import { storify } from '/@/libs/storify'
import toast from 'solid-toast'
import { WebsiteSettings } from '/@/components/WebsiteSettings'
import { InputLabel } from '/@/components/Input'
import { InputBar } from '/@/components/Input'
import { FormTextBig } from '/@/components/AppsNew'
import { PortPublicationSettings } from '/@/components/PortPublications'
import { userFromId, users } from '/@/libs/useAllUsers'
import { UserSearch } from '/@/components/UserSearch'
import useModal from '/@/libs/useModal'

const ContentContainer = styled('div', {
  base: {
    marginTop: '24px',
    display: 'grid',
    gridTemplateColumns: '380px 1fr',
    gap: '40px',
    position: 'relative',
  },
})
const SidebarContainer = styled('div', {
  base: {
    position: 'sticky',
    top: '64px',
    padding: '24px 40px',
    backgroundColor: vars.bg.white1,
    borderRadius: '4px',
    border: `1px solid ${vars.bg.white4}`,
  },
})
const SidebarOptions = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '12px',

    fontSize: '20px',
    color: vars.text.black1,
  },
})
const SidebarNavAnchor = styled('a', {
  base: {
    color: vars.text.black2,
    textDecoration: 'none',
    selectors: {
      '&:hover': {
        color: vars.text.black1,
      },
    },
  },
})
const ConfigsContainer = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '24px',
  },
})
const SettingFieldSet = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '16px',
    padding: '24px',
    border: `1px solid ${vars.bg.white4}`,
    borderRadius: '4px',
    background: vars.bg.white1,
  },
})

export default () => {
  const params = useParams()
  const [app, { refetch: refetchApp }] = createResource(
    () => params.id,
    (id) => client.getApplication({ id }),
  )
  const [repo] = createResource(
    () => app()?.repositoryId,
    (id) => client.getRepository({ repositoryId: id }),
  )
  const loaded = () => !!(app() && repo())

  const GeneralConfigsContainer: Component = () => {
    // 現在の設定で初期化
    const [generalConfig, setGeneralConfig] = createStore({
      name: app().name,
      refName: app().refName,
    })
    let formContainer: HTMLFormElement

    const updateGeneralSettings: JSX.EventHandler<HTMLButtonElement, MouseEvent> = async (e) => {
      // prevent default form submit (reload page)
      e.preventDefault()

      // validate form
      if (!formContainer.reportValidity()) {
        return
      }

      const updateApplicationRequest = new UpdateApplicationRequest({
        id: app().id,
        name: generalConfig.name,
        refName: generalConfig.refName,
      })

      try {
        await client.updateApplication(updateApplicationRequest)
        toast.success('アプリ設定を更新しました')
        refetchApp()
      } catch (e) {
        handleAPIError(e, 'アプリ設定の更新に失敗しました')
      }
    }

    return (
      <form ref={formContainer}>
        <SettingFieldSet>
          <FormTextBig id='general-settings'>General settings</FormTextBig>
          <div>
            <InputLabel>Application Name</InputLabel>
            <InputBar
              placeholder='my-app'
              value={generalConfig.name}
              onChange={(e) => setGeneralConfig('name', e.currentTarget.value)}
              required
            />
          </div>
          <div>
            <InputLabel>Branch Name</InputLabel>
            <InputBar
              placeholder='main'
              value={generalConfig.refName}
              onChange={(e) => setGeneralConfig('refName', e.currentTarget.value)}
              required
            />
          </div>
          <Button color='black1' size='large' onclick={updateGeneralSettings} type='submit'>
            Save
          </Button>
        </SettingFieldSet>
      </form>
    )
  }

  const BuildConfigsContainer: Component = () => {
    type BuildConfigMethod = ApplicationConfig['buildConfig']['case']
    const [runtimeConfig, setRuntimeConfig] = createStore<RuntimeConfig>(new RuntimeConfig())
    const [buildConfigMethod, setBuildConfigMethod] = createSignal<BuildConfigMethod>(app().config.buildConfig.case)
    const [buildConfig, setBuildConfig] = createStore<{
      [K in BuildConfigMethod]: Extract<ApplicationConfig['buildConfig'], { case: K }>
    }>({
      runtimeBuildpack: {
        case: 'runtimeBuildpack',
        value: storify(
          new BuildConfigRuntimeBuildpack({
            runtimeConfig: runtimeConfig,
          }),
        ),
      },
      runtimeCmd: {
        case: 'runtimeCmd',
        value: storify(
          new BuildConfigRuntimeCmd({
            runtimeConfig: runtimeConfig,
          }),
        ),
      },
      runtimeDockerfile: {
        case: 'runtimeDockerfile',
        value: storify(
          new BuildConfigRuntimeDockerfile({
            runtimeConfig: runtimeConfig,
          }),
        ),
      },
      staticCmd: {
        case: 'staticCmd',
        value: storify(new BuildConfigStaticCmd()),
      },
      staticDockerfile: {
        case: 'staticDockerfile',
        value: storify(new BuildConfigStaticDockerfile()),
      },
    })

    // 現在のビルド設定を反映
    onMount(() => {
      setBuildConfigMethod(app().config.buildConfig.case)
      setBuildConfig({
        [app().config.buildConfig.case]: {
          case: app().config.buildConfig.case,
          value: storify(app().config.buildConfig.value),
        },
      })
    })

    const updateBuildSettings: JSX.EventHandler<HTMLButtonElement, MouseEvent> = async (e) => {
      // prevent default form submit (reload page)
      e.preventDefault()

      const updateApplicationRequest = new UpdateApplicationRequest({
        id: app().id,
        config: {
          buildConfig: buildConfig[buildConfigMethod()],
        },
      })

      try {
        await client.updateApplication(updateApplicationRequest)
        toast.success('ビルド設定を更新しました')
        refetchApp()
      } catch (e) {
        handleAPIError(e, 'ビルド設定の更新に失敗しました')
      }
    }

    return (
      <SettingFieldSet>
        <FormTextBig id='build-settings'>Build Settings</FormTextBig>
        <BuildConfigs
          setBuildConfig={setBuildConfig}
          buildConfig={buildConfig}
          runtimeConfig={runtimeConfig}
          setRuntimeConfig={setRuntimeConfig}
          buildConfigMethod={buildConfigMethod()}
          setBuildConfigMethod={setBuildConfigMethod}
        />
        <Button color='black1' size='large' onclick={updateBuildSettings} type='submit'>
          Save
        </Button>
      </SettingFieldSet>
    )
  }

  const WebsitesConfigContainer: Component = () => {
    const [websites, setWebsites] = createStore<CreateWebsiteRequest[]>([])
    // 現在のウェブサイト設定を反映 (`onMount`ではrefetch時に反映されないので`createEffect`を使用)
    createEffect(() => {
      setWebsites(app().websites.map((website) => storify(website)))
    })

    const updateWebsites = async () => {
      const updateApplicationRequest = new UpdateApplicationRequest({
        id: app().id,
        websites: websites,
      })

      try {
        await client.updateApplication(updateApplicationRequest)
        toast.success('ウェブサイト設定を更新しました')
        refetchApp()
      } catch (e) {
        handleAPIError(e, 'ウェブサイト設定の更新に失敗しました')
      }
    }

    return (
      <SettingFieldSet>
        <FormTextBig id='website-settings'>Website Settings</FormTextBig>
        <WebsiteSettings websiteConfigs={websites} setWebsiteConfigs={setWebsites} />
        <Button color='black1' size='large' onclick={updateWebsites} type='submit'>
          Save
        </Button>
      </SettingFieldSet>
    )
  }

  const PortPublicationConfigContainer: Component = () => {
    const [currentPortPublications, setCurrentPortPublications] = createStore<PortPublication[]>([])
    // 現在のポート設定を反映
    onMount(() => {
      setCurrentPortPublications(app().portPublications.map((PortPublication) => storify(PortPublication)))
    })

    const updatePortPublications = async () => {
      const updateApplicationRequest = new UpdateApplicationRequest({
        id: app().id,
        portPublications: currentPortPublications,
      })

      try {
        await client.updateApplication(updateApplicationRequest)
        toast.success('ポート設定を更新しました')
        refetchApp()
      } catch (e) {
        handleAPIError(e, 'ポート設定の更新に失敗しました')
      }
    }

    return (
      <SettingFieldSet>
        <FormTextBig id='port-settings'>Port Publication Settings</FormTextBig>
        <PortPublicationSettings
          portPublications={currentPortPublications}
          setPortPublications={setCurrentPortPublications}
        />
        <Button color='black1' size='large' onclick={updatePortPublications} type='submit'>
          Save
        </Button>
      </SettingFieldSet>
    )
  }

  const OwnerConfigContainer: Component = () => {
    const { Modal, open } = useModal()

    const nonOwnerUsers = createMemo(() => {
      return users()?.filter((user) => !app().ownerIds.includes(user.id)) ?? []
    })

    const handleAddOwner = async (user: User) => {
      const updateApplicationRequest = new UpdateApplicationRequest({
        id: app().id,
        ownerIds: app().ownerIds.concat(user.id),
      })

      try {
        await client.updateApplication(updateApplicationRequest)
        toast.success('アプリオーナーを追加しました')
        refetchApp()
      } catch (e) {
        handleAPIError(e, 'アプリオーナーの追加に失敗しました')
      }
    }
    const handleDeleteOwner = async (owner: User) => {
      const updateApplicationRequest = new UpdateApplicationRequest({
        id: app().id,
        ownerIds: app().ownerIds.filter((id) => id !== owner.id),
      })

      try {
        await client.updateApplication(updateApplicationRequest)
        toast.success('アプリのオーナーを削除しました')
        refetchApp()
      } catch (e) {
        handleAPIError(e, 'アプリのオーナーの削除に失敗しました')
      }
    }

    return (
      <>
        <SettingFieldSet>
          <FormTextBig id='owner-settings'>Owner Settings</FormTextBig>
          <Button color='black1' size='large' onclick={open}>
            アプリオーナーを追加する
          </Button>
          <UserSearch users={app().ownerIds.map((userId) => userFromId(userId))}>
            {(user) => (
              <Button
                color='black1'
                size='large'
                onclick={() => {
                  handleDeleteOwner(user)
                }}
              >
                削除
              </Button>
            )}
          </UserSearch>
        </SettingFieldSet>
        <Modal>
          <UserSearch users={nonOwnerUsers()}>
            {(user) => (
              <Button
                color='black1'
                size='large'
                onclick={() => {
                  handleAddOwner(user)
                }}
              >
                追加
              </Button>
            )}
          </UserSearch>
        </Modal>
      </>
    )
  }

  return (
    <Container>
      <Header />
      <Show when={loaded()}>
        <AppNav repoName={repo().name} appName={app().name} appID={app().id} />
        <ContentContainer>
          <div>
            <SidebarContainer>
              <SidebarOptions>
                <SidebarNavAnchor href='#general-settings'>General</SidebarNavAnchor>
                <SidebarNavAnchor href='#build-settings'>Build</SidebarNavAnchor>
                <SidebarNavAnchor href='#website-settings'>Website</SidebarNavAnchor>
                <SidebarNavAnchor href='#port-settings'>Port Publication</SidebarNavAnchor>
                <SidebarNavAnchor href='#owner-settings'>Owner</SidebarNavAnchor>
              </SidebarOptions>
            </SidebarContainer>
          </div>
          <ConfigsContainer>
            <GeneralConfigsContainer />
            <BuildConfigsContainer />
            <WebsitesConfigContainer />
            <PortPublicationConfigContainer />
            <OwnerConfigContainer />
          </ConfigsContainer>
        </ContentContainer>
      </Show>
    </Container>
  )
}