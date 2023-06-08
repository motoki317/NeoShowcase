import { useParams } from '@solidjs/router'
import { Component, For, JSX, Show, createEffect, createMemo, createResource, createSignal, onMount } from 'solid-js'
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
  ApplicationEnvVar,
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
  Website,
} from '/@/api/neoshowcase/protobuf/gateway_pb'
import { BuildConfigs } from '/@/components/BuildConfigs'
import { storify } from '/@/libs/storify'
import toast from 'solid-toast'
import { WebsiteSetting } from '/@/components/WebsiteSettings'
import { InputLabel } from '/@/components/Input'
import { InputBar } from '/@/components/Input'
import { FormButton, FormTextBig } from '/@/components/AppsNew'
import { PortPublicationSettings } from '/@/components/PortPublications'
import { userFromId, users } from '/@/libs/useAllUsers'
import { UserSearch } from '/@/components/UserSearch'
import useModal from '/@/libs/useModal'
import { notEqual } from 'assert'

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
          <Button color='black1' size='large' width='auto' onclick={updateGeneralSettings} type='submit'>
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
        <Button color='black1' size='large' width='auto' onclick={updateBuildSettings} type='submit'>
          Save
        </Button>
      </SettingFieldSet>
    )
  }

  const WebsitesConfigContainer: Component = () => {
    // アプリにすでに存在するウェブサイト設定: `Website`
    // 新規追加するウェブサイト設定: `CreateWebsiteRequest`
    const [websites, setWebsites] = createStore<(Website | CreateWebsiteRequest)[]>([])
    // 現在のウェブサイト設定を反映 (`onMount`ではrefetch時に反映されないので`createEffect`を使用)
    createEffect(() => {
      setWebsites(app().websites.map((website) => storify(website)))
    })

    const AddWebsites = async (website: CreateWebsiteRequest) => {
      const updateApplicationRequest = new UpdateApplicationRequest({
        id: app().id,
        newWebsites: [website],
      })

      try {
        await client.updateApplication(updateApplicationRequest)
        toast.success('ウェブサイト設定を追加しました')
        refetchApp()
      } catch (e) {
        handleAPIError(e, 'ウェブサイト設定の追加に失敗しました')
      }
    }

    const deleteWebsites = async (websiteId: Website['id']) => {
      const updateApplicationRequest = new UpdateApplicationRequest({
        id: app().id,
        deleteWebsites: [{ id: websiteId }],
      })

      try {
        await client.updateApplication(updateApplicationRequest)
        toast.success('ウェブサイト設定を削除しました')
        refetchApp()
      } catch (e) {
        handleAPIError(e, 'ウェブサイト設定の削除に失敗しました')
      }
    }

    const updateWebsites = async (website: Website) => {
      const updateApplicationRequest = new UpdateApplicationRequest({
        id: app().id,
        deleteWebsites: [{ id: website.id }],
        newWebsites: [website],
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
        <For each={websites}>
          {(website, i) => (
            <WebsiteSetting
              website={website}
              deleteWebsite={() => {
                if (website instanceof CreateWebsiteRequest) {
                  // 新規に追加したウェブサイト設定は配列から削除する
                  setWebsites([...websites.slice(0, i()), ...websites.slice(i() + 1)])
                } else {
                  // もとからあるウェブサイト設定は削除リクエストを送る
                  deleteWebsites(website.id)
                }
              }}
              setWebsite={(key, value) => {
                setWebsites(i(), key, value)
              }}
              saveWebsite={() => {
                if (website instanceof CreateWebsiteRequest) {
                  // 新規に追加したウェブサイト設定は追加リクエストを送る
                  AddWebsites(website)
                } else {
                  // もとからあるウェブサイト設定は更新処理を実行する
                  updateWebsites(website)
                }
              }}
            />
          )}
        </For>
        <FormButton>
          <Button
            onclick={() => {
              setWebsites([...websites, storify(new CreateWebsiteRequest())])
            }}
            color='black1'
            size='large'
            width='auto'
            type='button'
          >
            Add website setting
          </Button>
        </FormButton>
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
        <Button color='black1' size='large' width='auto' onclick={updatePortPublications} type='submit'>
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
          <Button color='black1' size='large' width='auto' onclick={open}>
            アプリオーナーを追加する
          </Button>
          <UserSearch users={app().ownerIds.map((userId) => userFromId(userId))}>
            {(user) => (
              <Button
                color='black1'
                size='large'
                width='auto'
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
                width='auto'
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

  const EnvVarConfigContainer: Component = () => {
    const [envVars] = createResource(
      () => app().id,
      (id) => client.getEnvVars({ id }),
    )

    const EditEnvVarContainer: Component<{
      envVar: ApplicationEnvVar
    }> = (props) => {
      const [isEditing, setIsEditing] = createSignal(false)
      let formRef: HTMLFormElement
      let keyInputRef: HTMLInputElement
      let valueInputRef: HTMLInputElement

      const handleUpdateEnvVar: JSX.EventHandler<HTMLButtonElement, MouseEvent> = async (e) => {
        // prevent default form submit (reload page)
        e.preventDefault()

        // validate form
        if (!formRef.reportValidity()) {
          return
        }

        try {
          // TODO: key変更時は旧keyの環境変数を削除する
          await client.setEnvVar({
            applicationId: app().id,
            key: keyInputRef.value,
            value: valueInputRef.value,
          })
          toast.success('環境変数を更新しました')
          refetchApp()
          setIsEditing(false)
        } catch (e) {
          handleAPIError(e, '環境変数の更新に失敗しました')
        }
      }

      const handleDeleteEnvVar: JSX.EventHandler<HTMLButtonElement, MouseEvent> = async (e) => {
        // prevent default form submit (reload page)
        e.preventDefault()

        try {
          // TODO: 環境変数を削除する
          throw new Error('not implemented')
          // toast.success('環境変数を削除しました')
          // refetchApp()
          // setIsEditing(false)
        } catch (e) {
          handleAPIError(e, '環境変数の削除に失敗しました')
        }
      }

      return (
        <form ref={formRef}>
          <InputBar
            type='text'
            disabled={!isEditing()}
            required
            placeholder='KEY'
            ref={keyInputRef}
            value={props.envVar.key}
          />
          <InputBar
            type='text'
            disabled={!isEditing()}
            required
            placeholder='VALUE'
            ref={valueInputRef}
            value={props.envVar.value}
          />
          <Show
            when={!isEditing()}
            fallback={
              <Button color='black1' size='large' width='auto' type='submit' onclick={handleUpdateEnvVar}>
                Save
              </Button>
            }
          >
            <Button color='black1' size='large' width='auto' type='button' onclick={() => setIsEditing(true)}>
              Edit
            </Button>
            <Button color='black1' size='large' width='auto' type='button' onclick={handleDeleteEnvVar}>
              Delete
            </Button>
          </Show>
        </form>
      )
    }

    const AddEnvVarContainer: Component = () => {
      let formRef: HTMLFormElement
      let keyInputRef: HTMLInputElement
      let valueInputRef: HTMLInputElement

      const handleAddEnvVar: JSX.EventHandler<HTMLButtonElement, MouseEvent> = async (e) => {
        // prevent default form submit (reload page)
        e.preventDefault()

        // validate form
        if (!formRef.reportValidity()) {
          return
        }

        try {
          await client.setEnvVar({
            applicationId: app().id,
            key: keyInputRef.value,
            value: valueInputRef.value,
          })
          toast.success('環境変数を追加しました')
          refetchApp()
        } catch (e) {
          handleAPIError(e, '環境変数の追加に失敗しました')
        }
      }

      return (
        <form ref={formRef}>
          <InputBar type='text' required placeholder='KEY' ref={keyInputRef} />
          <InputBar type='text' required placeholder='VALUE' ref={valueInputRef} />
          <Button color='black1' size='large' width='auto' type='submit' onclick={handleAddEnvVar}>
            Add
          </Button>
        </form>
      )
    }

    return (
      <SettingFieldSet>
        <FormTextBig id='env-var-settings'>Environment Variable Settings</FormTextBig>
        <div>
          <For each={envVars()?.variables}>
            {(envVar) => {
              return <EditEnvVarContainer envVar={envVar} />
            }}
          </For>
          <AddEnvVarContainer />
        </div>
      </SettingFieldSet>
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
                <SidebarNavAnchor href='#env-var-settings'>Environment Variable</SidebarNavAnchor>
              </SidebarOptions>
            </SidebarContainer>
          </div>
          <ConfigsContainer>
            <GeneralConfigsContainer />
            <BuildConfigsContainer />
            <WebsitesConfigContainer />
            <PortPublicationConfigContainer />
            <OwnerConfigContainer />
            <EnvVarConfigContainer />
          </ConfigsContainer>
        </ContentContainer>
      </Show>
    </Container>
  )
}
