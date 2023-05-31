import { Header } from '/@/components/Header'
import { createResource, createSignal, JSX, For, JSXElement, Switch, Match } from 'solid-js'
import { Radio, RadioItem } from '/@/components/Radio'
import { client } from '/@/libs/api'
import {
  Application,
  Repository,
  AuthenticationType,
  PortPublicationProtocol,
  CreateApplicationRequest,
  ApplicationConfig,
  BuildConfigRuntimeBuildpack,
  CreateWebsiteRequest,
  PortPublication,
  BuildConfigRuntimeCmd,
  BuildConfigRuntimeDockerfile,
  BuildConfigStaticCmd,
  BuildConfigStaticDockerfile,
  RuntimeConfig,
} from '/@/api/neoshowcase/protobuf/gateway_pb'
import { A, useNavigate, useSearchParams } from '@solidjs/router'
import { BsArrowLeftShort } from 'solid-icons/bs'
import { Container } from '/@/libs/layout'
import { vars } from '/@/theme'
import { styled } from '@macaron-css/solid'
import { Button } from '/@/components/Button'
import { Checkbox } from '/@/components/Checkbox'
import { providerToIcon, repositoryURLToProvider } from '/@/libs/application'
import { createStore, SetStoreFunction } from 'solid-js/store'
import { Empty } from '@bufbuild/protobuf'
import toast from 'solid-toast'
import { ConnectError } from '@bufbuild/connect'

const [repos] = createResource(() => client.getRepositories({}))
const [apps] = createResource(() => client.getApplications({}))

const loaded = () => !!(repos() && apps())

const buildConfigItems: RadioItem[] = [
  { value: 'runtimeBuildpack', title: 'runtime buildpack' },
  { value: 'runtimeCmd', title: 'runtime cmd' },
  { value: 'runtimeDockerfile', title: 'runtime dockerfile' },
  { value: 'staticCmd', title: 'static cmd' },
  { value: 'staticDockerfile', title: 'static dockerfile' },
]

const AppTitle = styled('div', {
  base: {
    marginTop: '48px',
    height: '46px',
    lineHeight: '46px',

    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
  },
})

const AppsTitle = styled('div', {
  base: {
    fontSize: '32px',
    fontWeight: 700,
    color: vars.text.black1,
    display: 'flex',
  },
})

const Arrow = styled('div', {
  base: {
    fontSize: '32px',
    color: vars.text.black1,
    display: 'flex',
  },
})

const ContentContainer = styled('div', {
  base: {
    marginTop: '24px',
    display: 'grid',
    gap: '40px',
  },
})

const MainContentContainer = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '20px',
  },
})

const FormContainer = styled('form', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '20px',

    background: vars.bg.white3,
    border: `1px solid ${vars.bg.white4}`,
    borderRadius: '4px',
    padding: '8px 12px',
  },
})

const Form = styled('div', {
  base: {},
})

const FormButton = styled('div', {
  base: {
    marginLeft: '8px',
  },
})

const InputLabel = styled('div', {
  base: {
    fontSize: '16px',
    alignItems: 'center',
    fontWeight: 700,
    color: vars.text.black1,

    marginBottom: '4px',
  },
})

const FormTextBig = styled('div', {
  base: {
    fontSize: '20px',
    alignItems: 'center',
    fontWeight: 900,
    color: vars.text.black1,

    marginBottom: '4px',
  },
})

const InputBar = styled('input', {
  base: {
    padding: '8px 12px',
    borderRadius: '4px',
    border: `1px solid ${vars.bg.white4}`,
    fontSize: '14px',
    marginLeft: '4px',

    width: '320px',

    display: 'flex',
    flexDirection: 'column',

    '::placeholder': {
      color: vars.text.black3,
    },
  },
})

const FormCheckBox = styled('div', {
  base: {
    background: vars.bg.white1,
    border: `1px solid ${vars.bg.white4}`,
    borderRadius: '4px',
    marginLeft: '4px',
    padding: '8px 12px',

    width: '320px',
  },
})

const FormWebsite = styled('div', {
  base: {
    background: vars.bg.white2,
    border: `1px solid ${vars.bg.white4}`,
    borderRadius: '4px',
    marginLeft: '4px',
    padding: '8px 12px',

    display: 'flex',
    flexDirection: 'column',
    gap: '12px',
  },
})

const FormWebsiteButton = styled('div', {
  base: {
    display: 'flex',
    gap: '8px',
    marginBottom: '4px',
  },
})

const FormRadio = styled('div', {
  base: {
    background: vars.bg.white2,
    border: `1px solid ${vars.bg.white4}`,
    borderRadius: '4px',
    marginLeft: '4px',
    padding: '8px 12px',

    display: 'flex',
    flexDirection: 'column',
    gap: '12px',
  },
})

const RepositoriesContainer = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: 'px',
  },
})

const SettingsContainer = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '10px',
  },
})

interface Website {
  fqdn: string
  path_prefix: string
  strip_prefix: boolean
  https: boolean
  h2c: boolean
  http_port: number
  authenticationType: number
}

const EmptyWebsite: Website = {
  fqdn: '',
  path_prefix: '',
  strip_prefix: false,
  https: false,
  h2c: false,
  http_port: 0,
  authenticationType: AuthenticationType.OFF,
}

interface WebsiteProps {
  website: CreateWebsiteRequest
  setWebsite: (valueName, value) => void
  deleteWebsite: () => void
}

const Website = (props: WebsiteProps) => {
  return (
    <FormWebsite>
      <Form>
        <InputLabel>ドメイン名</InputLabel>
        <InputBar placeholder='example.ns.trap.jp' />
      </Form>
      <Form>
        <InputLabel>Path Prefix</InputLabel>
        <InputBar placeholder='/' />
      </Form>
      <Form>
        <FormCheckBox>
          <Checkbox
            selected={props.website.stripPrefix}
            setSelected={(selected) => props.setWebsite('stripPrefix', selected)}
          >
            Strip Path Prefix
          </Checkbox>
          <Checkbox selected={props.website.https} setSelected={(selected) => props.setWebsite('https', selected)}>
            https
          </Checkbox>
          <Checkbox selected={props.website.h2c} setSelected={(selected) => props.setWebsite('h2c', selected)}>
            (advanced) アプリ通信にh2cを用いる
          </Checkbox>
        </FormCheckBox>
      </Form>
      <Form>
        <InputLabel>アプリのHTTP Port番号</InputLabel>
        <InputBar placeholder='80' />
      </Form>
      <Form>
        <Radio
          items={authenticationTypeItems}
          selected={props.website.authentication}
          setSelected={(selected) => props.setWebsite('authentication', selected)}
        />
      </Form>
      <FormWebsiteButton>
        <Button onclick={props.deleteWebsite} color='black1' size='large' type='button'>
          Delete website setting
        </Button>
      </FormWebsiteButton>
    </FormWebsite>
  )
}

const authenticationTypeItems: RadioItem[] = [
  { value: AuthenticationType.OFF, title: 'OFF' },
  { value: AuthenticationType.SOFT, title: 'SOFT' },
  { value: AuthenticationType.HARD, title: 'HARD' },
]

interface Ports {
  internet_port: number
  application_port: number
  protocol: number
}

const EmptyPortPublication: Ports = {
  internet_port: 0,
  application_port: 0,
  protocol: PortPublicationProtocol.TCP,
}

interface PortPublicationProps {
  portPublication: Ports
  setPortPublication: (Ports) => void
  deletePortPublication: () => void
}

const PortPublications = (props: PortPublicationProps) => {
  return (
    <FormWebsite>
      <Form>
        <InputLabel>Internet Port</InputLabel>
        <InputBar placeholder='30000' />
      </Form>
      <Form>
        <InputLabel>Application Port</InputLabel>
        <InputBar placeholder='30001' />
      </Form>
      <Form>
        <Radio
          items={protocolItems}
          selected={props.portPublication.protocol}
          setSelected={(proto) => props.setPortPublication({ ...props.portPublication, protocol: proto })}
        />
      </Form>
      <FormWebsiteButton>
        <Button onclick={props.deletePortPublication} color='black1' size='large' type='button'>
          Delete port publication
        </Button>
      </FormWebsiteButton>
    </FormWebsite>
  )
}

const protocolItems: RadioItem[] = [
  { value: 0, title: 'TCP' },
  { value: 1, title: 'UDP' },
]

export interface FormRuntimeConfigProps {
  runtimeConfig: RuntimeConfig
  setRuntimeConfig: SetStoreFunction<RuntimeConfig>
}

const FormRuntimeConfig = (props: FormRuntimeConfigProps) => {
  return (
    <>
      <Form>
        <InputLabel>Database (使うデーターベースにチェック)</InputLabel>
        <FormCheckBox>
          <Checkbox
            selected={props.runtimeConfig.useMariadb}
            setSelected={(useMariadb) => props.setRuntimeConfig('useMariadb', useMariadb)}
          >
            MariaDB
          </Checkbox>
          <Checkbox
            selected={props.runtimeConfig.useMongodb}
            setSelected={(useMongodb) => props.setRuntimeConfig('useMongodb', useMongodb)}
          >
            MongoDB
          </Checkbox>
        </FormCheckBox>
      </Form>
      <Form>
        <InputLabel>Entrypoint</InputLabel>
        <InputBar
          placeholder=''
          value={props.runtimeConfig.entrypoint}
          onInput={(e) => props.setRuntimeConfig('entrypoint', e.target.value)}
        />
      </Form>
      <Form>
        <InputLabel>Command</InputLabel>
        <InputBar
          placeholder=''
          value={props.runtimeConfig.command}
          onInput={(e) => props.setRuntimeConfig('command', e.target.value)}
        />
      </Form>
    </>
  )
}

const RepositoryInfoContainer = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    gap: '10px',
  },
})

const RepoName = styled('div', {
  base: {
    fontSize: '16px',
    fontWeight: 500,
    color: vars.text.black1,
  },
})

export interface RepositoryInfoProps {
  repo: Repository
  apps: Application[]
}

const RepositoryInfoBackground = styled('div', {
  base: {
    display: 'flex',

    background: vars.bg.white3,
    border: `1px solid ${vars.bg.white4}`,
    borderRadius: '4px',
    padding: '8px 12px',
  },
})

const SmallText = styled('div', {
  base: {
    display: 'flex',
    fontSize: '11px',
    color: vars.text.black3,
  },
})

const RepositoryInfo = (props: RepositoryInfoProps): JSXElement => {
  const provider = repositoryURLToProvider(props.repo.url)
  return (
    <RepositoryInfoBackground>
      <RepositoryInfoContainer>
        {providerToIcon(provider)}
        <RepoName>{props.repo.name}</RepoName>
        <SmallText>{props.repo.url}</SmallText>
      </RepositoryInfoContainer>
    </RepositoryInfoBackground>
  )
}

export default () => {
  const navigate = useNavigate()
  const appsByRepo = () =>
    loaded() &&
    apps().applications.reduce((acc, app) => {
      if (!acc[app.repositoryId]) acc[app.repositoryId] = []
      acc[app.repositoryId].push(app)
      return acc
    }, {} as Record<string, Application[]>)

  const [websites, setWebsites] = createSignal<Website[]>([])
  const [portPublications, setPortPublications] = createSignal<Ports[]>([])

  const [checkBoxBuildCmdShell, setCheckBoxBuildCmdShell] = createSignal(false)

  const [createApplicationRequest, setCreateApplicationRequest] = createStore(
    new CreateApplicationRequest({
      config: new ApplicationConfig(),
      websites: [],
      portPublications: [],
    }),
  )

  const [websiteConfigs, setWebsiteConfigs] = createStore<CreateWebsiteRequest[]>([])
  const [fieldsPortPublication, setFieldsPortPublication] = createStore<PortPublication[]>([])

  // Build Config
  type BuildConfigMethod = ApplicationConfig['buildConfig']['case']
  const [runtimeConfig, setRuntimeConfig] = createStore<RuntimeConfig>(new RuntimeConfig())
  const [buildConfigMethod, setBuildConfigMethod] = createSignal<BuildConfigMethod>('runtimeBuildpack')
  const [buildConfig, setBuildConfig] = createStore<{
    [K in BuildConfigMethod]: Extract<ApplicationConfig['buildConfig'], { case: K }>
  }>({
    runtimeBuildpack: {
      case: 'runtimeBuildpack',
      value: new BuildConfigRuntimeBuildpack({
        runtimeConfig: new RuntimeConfig(),
      }),
    },
    runtimeCmd: {
      case: 'runtimeCmd',
      value: new BuildConfigRuntimeCmd({
        runtimeConfig: new RuntimeConfig(),
      }),
    },
    runtimeDockerfile: {
      case: 'runtimeDockerfile',
      value: new BuildConfigRuntimeDockerfile({
        runtimeConfig: new RuntimeConfig(),
      }),
    },
    staticCmd: {
      case: 'staticCmd',
      value: new BuildConfigStaticCmd(),
    },
    staticDockerfile: {
      case: 'staticDockerfile',
      value: new BuildConfigStaticDockerfile(),
    },
  })

  const [searchParams] = useSearchParams()
  setCreateApplicationRequest('repositoryId', searchParams.repositoryID)

  let formContainer: HTMLFormElement

  const createApplication: JSX.EventHandler<HTMLInputElement, MouseEvent> = async (e) => {
    // prevent default form submit (reload page)
    e.preventDefault()

    // validate form
    if (formContainer.reportValidity()) {
      setCreateApplicationRequest('config', 'buildConfig', buildConfig[buildConfigMethod()])
      try {
        const res = await client.createApplication(createApplicationRequest)
        toast.success('アプリケーションを登録しました')
        // Application詳細ページに遷移
        navigate(`/apps/${res.id}`)
      } catch (e) {
        console.error(e)
        // gRPCエラー
        if (e instanceof ConnectError) {
          toast.error('アプリケーションの登録に失敗しました\n' + e.message)
        }
      }
    }
  }

  const SelectRepository = (): JSX.Element => {
    return (
      <>
        <ContentContainer>
          <MainContentContainer>
            {loaded() &&
              repos()
                .repositories.filter((r) => r.id === searchParams.repositoryID)
                .map((r) => <RepositoryInfo repo={r} apps={appsByRepo()[r.id] || []} />)}

            <FormContainer ref={formContainer}>
              <Form>
                <InputLabel>Application Name</InputLabel>
                <InputBar
                  placeholder=''
                  value={createApplicationRequest.name}
                  onInput={(e) => setCreateApplicationRequest('name', e.target.value)}
                />
              </Form>

              <Form>
                <InputLabel>Branch Name</InputLabel>
                <InputBar
                  placeholder='master'
                  value={createApplicationRequest.refName}
                  onInput={(e) => setCreateApplicationRequest('refName', e.target.value)}
                />
              </Form>

              <Form>
                <FormTextBig>Build Config</FormTextBig>
                <FormRadio>
                  <Form>
                    <Radio items={buildConfigItems} selected={buildConfigMethod()} setSelected={setBuildConfigMethod} />
                  </Form>

                  <Switch>
                    <Match when={buildConfigMethod() === 'runtimeBuildpack'}>
                      <FormRuntimeConfig runtimeConfig={runtimeConfig} setRuntimeConfig={setRuntimeConfig} />
                      <Form>
                        <InputLabel>Context</InputLabel>
                        <InputBar
                          placeholder=''
                          value={buildConfig.runtimeBuildpack.value.context}
                          onInput={(e) => setBuildConfig('runtimeBuildpack', 'value', 'context', e.target.value)}
                        />
                      </Form>
                    </Match>

                    <Match when={buildConfigMethod() === 'runtimeCmd'}>
                      <FormRuntimeConfig runtimeConfig={runtimeConfig} setRuntimeConfig={setRuntimeConfig} />
                      <Form>
                        <InputLabel>Base image</InputLabel>
                        <InputBar
                          placeholder=''
                          value={buildConfig.runtimeCmd.value.baseImage}
                          onInput={(e) => setBuildConfig('runtimeCmd', 'value', 'baseImage', e.target.value)}
                        />
                      </Form>
                      <Form>
                        <InputLabel>Build cmd</InputLabel>
                        <InputBar
                          placeholder=''
                          value={buildConfig.runtimeCmd.value.buildCmd}
                          onInput={(e) => setBuildConfig('runtimeCmd', 'value', 'buildCmd', e.target.value)}
                        />
                      </Form>
                      <Form>
                        <InputLabel>Build cmd shell</InputLabel>
                        <FormCheckBox>
                          <Checkbox
                            selected={buildConfig.runtimeCmd.value.buildCmdShell}
                            setSelected={(selected) => setBuildConfig('runtimeCmd', 'value', 'buildCmdShell', selected)}
                          >
                            Run build cmd with shell
                          </Checkbox>
                        </FormCheckBox>
                      </Form>
                    </Match>

                    <Match when={buildConfigMethod() === 'runtimeDockerfile'}>
                      <FormRuntimeConfig runtimeConfig={runtimeConfig} setRuntimeConfig={setRuntimeConfig} />
                      <Form>
                        <InputLabel>Dockerfile name</InputLabel>
                        <InputBar
                          placeholder=''
                          value={buildConfig.runtimeDockerfile.value.dockerfileName}
                          onInput={(e) =>
                            setBuildConfig('runtimeDockerfile', 'value', 'dockerfileName', e.target.value)
                          }
                        />
                      </Form>
                      <Form>
                        <InputLabel>Context</InputLabel>
                        <InputBar
                          placeholder=''
                          value={buildConfig.runtimeDockerfile.value.context}
                          onInput={(e) => setBuildConfig('runtimeDockerfile', 'value', 'context', e.target.value)}
                        />
                      </Form>
                    </Match>

                    <Match when={buildConfigMethod() === 'staticCmd'}>
                      <Form>
                        <InputLabel>Base image</InputLabel>
                        <InputBar
                          placeholder=''
                          value={buildConfig.staticCmd.value.baseImage}
                          onInput={(e) => setBuildConfig('staticCmd', 'value', 'baseImage', e.target.value)}
                        />
                      </Form>
                      <Form>
                        <InputLabel>Build cmd</InputLabel>
                        <InputBar
                          placeholder=''
                          value={buildConfig.staticCmd.value.buildCmd}
                          onInput={(e) => setBuildConfig('staticCmd', 'value', 'buildCmd', e.target.value)}
                        />
                      </Form>
                      <Form>
                        <InputLabel>Build cmd shell</InputLabel>
                        <FormCheckBox>
                          <Checkbox
                            selected={buildConfig.staticCmd.value.buildCmdShell}
                            setSelected={(selected) => setBuildConfig('staticCmd', 'value', 'buildCmdShell', selected)}
                          >
                            Run build cmd with shell
                          </Checkbox>
                        </FormCheckBox>
                      </Form>
                      <Form>
                        <InputLabel>Artifact path</InputLabel>
                        <InputBar
                          placeholder=''
                          value={buildConfig.staticCmd.value.artifactPath}
                          onInput={(e) => setBuildConfig('staticCmd', 'value', 'artifactPath', e.target.value)}
                        />
                      </Form>
                    </Match>

                    <Match when={buildConfigMethod() === 'staticDockerfile'}>
                      <Form>
                        <InputLabel>Dockerfile name</InputLabel>
                        <InputBar
                          placeholder=''
                          value={buildConfig.staticDockerfile.value.dockerfileName}
                          onInput={(e) => setBuildConfig('staticDockerfile', 'value', 'dockerfileName', e.target.value)}
                        />
                      </Form>
                      <Form>
                        <InputLabel>Context</InputLabel>
                        <InputBar
                          placeholder=''
                          value={buildConfig.staticDockerfile.value.context}
                          onInput={(e) => setBuildConfig('staticDockerfile', 'value', 'context', e.target.value)}
                        />
                      </Form>
                      <Form>
                        <InputLabel>Artifact path</InputLabel>
                        <InputBar
                          placeholder=''
                          value={buildConfig.staticDockerfile.value.artifactPath}
                          onInput={(e) => setBuildConfig('staticDockerfile', 'value', 'artifactPath', e.target.value)}
                        />
                      </Form>
                    </Match>
                  </Switch>
                </FormRadio>
              </Form>

              <Form>
                <FormTextBig>Website Setting</FormTextBig>
                <SettingsContainer>
                  <For each={websiteConfigs}>
                    {(website, i) => (
                      <Website
                        website={website}
                        setWebsite={(valueName, value) =>
                          setWebsiteConfigs((prev) => {
                            const newWebsites = [...prev]
                            newWebsites[i()][valueName] = value
                            return newWebsites
                          })
                        }
                        deleteWebsite={() =>
                          setWebsiteConfigs((current) => [...current.slice(0, i()), ...current.slice(i() + 1)])
                        }
                      />
                    )}
                  </For>

                  <FormButton>
                    <Button
                      onclick={() => {
                        setWebsiteConfigs([...websiteConfigs, new CreateWebsiteRequest()])
                      }}
                      color='black1'
                      size='large'
                      type='button'
                    >
                      Add website setting
                    </Button>
                  </FormButton>
                </SettingsContainer>
              </Form>

              <Form>
                <FormTextBig>Port Publication Setting</FormTextBig>
                <SettingsContainer>
                  <For each={portPublications()}>
                    {(portPublication, i) => (
                      <PortPublications
                        portPublication={portPublication}
                        setPortPublication={(portPublication) =>
                          setPortPublications((current) => [
                            ...current.slice(0, i()),
                            portPublication,
                            ...current.slice(i() + 1),
                          ])
                        }
                        deletePortPublication={() =>
                          setPortPublications((current) => [...current.slice(0, i()), ...current.slice(i() + 1)])
                        }
                      />
                    )}
                  </For>

                  <FormButton>
                    <Button
                      onclick={() => {
                        setPortPublications([...portPublications(), EmptyPortPublication])
                      }}
                      color='black1'
                      size='large'
                      type='button'
                    >
                      Add port publication
                    </Button>
                  </FormButton>
                </SettingsContainer>
              </Form>

              <Form>
                <InputLabel>Start on create</InputLabel>
                <FormCheckBox>
                  <Checkbox
                    selected={createApplicationRequest.startOnCreate}
                    setSelected={(selected) => setCreateApplicationRequest('startOnCreate', selected)}
                  >
                    start_on_create
                  </Checkbox>
                </FormCheckBox>
              </Form>

              <Button color='black1' size='large' onclick={createApplication} type="submit">
                + Create new Application
              </Button>

              <Button
                onclick={() => {
                  console.log(websiteConfigs)
                }}
                color='black1'
                size='large'
                type='button'
              >
                Debug
              </Button>
            </FormContainer>
          </MainContentContainer>
        </ContentContainer>
      </>
    )
  }

  return (
    <Container>
      <Header />
      <AppTitle>
        <A href={'/apps'}>
          <Arrow>
            <BsArrowLeftShort />
          </Arrow>
        </A>
        <AppsTitle>Create Application</AppsTitle>
      </AppTitle>
      <SelectRepository />
    </Container>
  )
}
