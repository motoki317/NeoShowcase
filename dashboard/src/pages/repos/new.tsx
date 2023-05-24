import { JSX, JSXElement, Match, Show, Switch, createEffect, createResource, createSignal } from 'solid-js'
import { createStore } from 'solid-js/store'
import { Empty } from '@bufbuild/protobuf'
import { styled } from '@macaron-css/solid'
import {
  CreateRepositoryAuth,
  CreateRepositoryAuthBasic,
  CreateRepositoryAuthSSH,
  CreateRepositoryRequest,
} from '/@/api/neoshowcase/protobuf/gateway_pb'
import { Button } from '/@/components/Button'
import { Header } from '/@/components/Header'
import { Radio } from '/@/components/Radio'
import { client } from '/@/libs/api'
import { Container } from '/@/libs/layout'
import { vars } from '/@/theme'

// copy from /pages/apps AppsTitle component
const PageTitle = styled('div', {
  base: {
    marginTop: '48px',
    fontSize: '32px',
    fontWeight: 'bold',
    color: vars.text.black1,
  },
})

// copy from /pages/apps
// and delete unnecessary styles
const ContentContainer = styled('div', {
  base: {
    marginTop: '24px',
  },
})

// copy from /pages/apps/new
const InputFormContainer = styled('div', {
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
const InputForm = styled('div', {
  base: {},
})
const InputFormText = styled('div', {
  base: {
    fontSize: '16px',
    alignItems: 'center',
    fontWeight: 700,
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

interface FormProps {
  label: string
  type?: JSX.InputHTMLAttributes<HTMLInputElement>['type']
  placeholder?: JSX.InputHTMLAttributes<HTMLInputElement>['placeholder']
  value: JSX.InputHTMLAttributes<HTMLInputElement>['value']
  onInput: JSX.InputHTMLAttributes<HTMLInputElement>['onInput']
}

const Form = (props: FormProps): JSXElement => {
  return (
    <InputForm>
      <InputFormText>{props.label}</InputFormText>
      <InputBar
        type={props.type ?? 'text'}
        placeholder={props.placeholder ?? ''}
        value={props.value}
        onInput={props.onInput}
      />
    </InputForm>
  )
}

const SshDetails = styled('div', {
  base: {
    color: vars.text.black2,
    marginBottom: '4px',
  },
})

const PublicKeyCode = styled('code', {
  base: {
    display: 'block',
    padding: '8px 12px',
    fontFamily: 'monospace',
    fontSize: '14px',
    background: vars.bg.white2,
    color: vars.text.black1,
    border: `1px solid ${vars.bg.white4}`,
    borderRadius: '4px',
  },
})

export default () => {
  // 認証方法 ("none" | "ssh" | "basic")
  type AuthMethod = CreateRepositoryAuth['auth']['case']
  const [authMethod, setAuthMethod] = createSignal<AuthMethod>('none')

  const [sshAuthConfig, setSshAuthConfig] = createStore(new CreateRepositoryAuthSSH())
  const [basicAuthConfig, setBasicAuthConfig] = createStore(new CreateRepositoryAuthBasic())

  const [requestConfig, setRequestConfig] = createStore(
    new CreateRepositoryRequest({
      auth: new CreateRepositoryAuth(),
    }),
  )

  const createRepository = async () => {
    // 認証方法に応じて認証情報を設定
    switch (authMethod()) {
      case 'none':
        setRequestConfig('auth', 'auth', { value: new Empty(), case: 'none' })
        break
      case 'ssh':
        setRequestConfig('auth', 'auth', { value: sshAuthConfig, case: 'ssh' })
        break
      case 'basic':
        setRequestConfig('auth', 'auth', { value: basicAuthConfig, case: 'basic' })
        break
    }

    const res = await client.createRepository(requestConfig)
    // TODO: navigate to repository page when success / show error message when failed
  }

  // URLからリポジトリ名を自動入力
  createEffect(() => {
    const segments = requestConfig.url.split('/')
    const lastSegment = segments.pop() || segments.pop() // 末尾のスラッシュを除去
    const repositoryName = lastSegment?.replace(/\.git$/, '') ?? ''
    setRequestConfig('name', repositoryName)
  })

  const [systemPublicKey] = createResource(() => client.getSystemPublicKey({}))

  return (
    <Container>
      <Header />
      <PageTitle>Create Repository</PageTitle>
      <ContentContainer>
        <InputFormContainer>
          <Form
            label='URL'
            type='url'
            placeholder='https://example.com/my-app.git'
            value={requestConfig.url}
            onInput={(e) => setRequestConfig('url', e.currentTarget.value)}
          />
          <Form
            label='リポジトリ名'
            placeholder='my-app'
            value={requestConfig.name}
            onInput={(e) => setRequestConfig('name', e.currentTarget.value)}
          />
          <InputForm>
            <InputFormText>認証方法</InputFormText>
            <Radio
              items={[
                { title: '認証を使用しない', value: 'none' },
                { title: 'Basic認証を使用', value: 'basic' },
                { title: 'SSH認証を使用', value: 'ssh' },
              ]}
              selected={authMethod()}
              setSelected={setAuthMethod}
            />
          </InputForm>
          <Switch>
            <Match when={authMethod() === 'basic'}>
              <Form
                label='ユーザー名'
                value={basicAuthConfig.username}
                onInput={(e) => setBasicAuthConfig('username', e.currentTarget.value)}
              />
              <Form
                label='パスワード'
                type='password'
                value={basicAuthConfig.password}
                onInput={(e) => setBasicAuthConfig('password', e.currentTarget.value)}
              />
            </Match>
            <Match when={authMethod() === 'ssh'}>
              <Form
                label='SSH秘密鍵'
                placeholder=''
                value={sshAuthConfig.sshKey}
                onInput={(e) => setSshAuthConfig('sshKey', e.currentTarget.value)}
              />
              <Show when={sshAuthConfig.sshKey.length === 0}>
                <div>
                  <SshDetails>
                    秘密鍵を入力せずにSSH認証でリポジトリを登録する場合、以下のSSH公開鍵が認証に使用されます。
                  </SshDetails>
                  <Switch>
                    <Match when={systemPublicKey.loading}>
                      <div>Loading...</div>
                    </Match>
                    <Match when={systemPublicKey()}>
                      <PublicKeyCode>{systemPublicKey().publicKey}</PublicKeyCode>
                    </Match>
                  </Switch>
                </div>
              </Show>
            </Match>
          </Switch>
          <Button color='black1' size='large' onclick={createRepository}>
            + Create new Repository
          </Button>
        </InputFormContainer>
      </ContentContainer>
    </Container>
  )
}
