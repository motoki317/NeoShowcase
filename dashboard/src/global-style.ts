import { vars } from '/@/theme'
import { globalStyle } from '@macaron-css/core'

globalStyle('*', {
  fontFamily: 'Noto Sans JP',
})

globalStyle('a', {
  textDecoration: 'none',
})

globalStyle('body', {
  margin: '0',
  backgroundColor: vars.bg.white2,
})