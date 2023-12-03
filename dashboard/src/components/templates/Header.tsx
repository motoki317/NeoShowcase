import { styled } from '@macaron-css/solid'
import { A } from '@solidjs/router'
import { Component, Show } from 'solid-js'
import LogoImage from '/@/assets/logo.svg?url'
import SmallLogoImage from '/@/assets/logo_small.svg?url'
import { user } from '/@/libs/api'
import { colorVars } from '/@/theme'
import { Button } from '../UI/Button'
import { UserMenuButton } from '../UI/UserMenuButton'

const Container = styled('div', {
  base: {
    width: '100%',
    height: '64px',
    padding: '10px 24px',
    flexShrink: 0,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-start',
    gap: '24px',
    borderBottom: `1px solid ${colorVars.semantic.ui.border}`,
  },
})

export const Header: Component = () => {
  return (
    <Container>
      <A href="/">
        {/* 画面幅が768px以下の時はSmallLogoImageを表示する */}
        <source srcset={SmallLogoImage} media="(max-width: 768px)" />
        <img src={LogoImage} alt="NeoShowcase logo" />
      </A>
      <A href="/apps">
        <Button size="medium" variants="text">
          Apps
        </Button>
      </A>
      <A href="/builds" style={{ 'margin-right': 'auto' }}>
        <Button size="medium" variants="text">
          Queue
        </Button>
      </A>
      <Show when={user()}>{(user) => <UserMenuButton user={user()} />}</Show>
    </Container>
  )
}
