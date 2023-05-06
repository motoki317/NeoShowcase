import { styled } from '@macaron-css/solid'
import { applicationState, ApplicationState } from '/@/libs/application'
import { JSX } from 'solid-js'
import { StatusIcon } from '/@/components/StatusIcon'
import { Application } from '/@/api/neoshowcase/protobuf/gateway_pb'

const Container = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    width: '100%',
  },
})

const ContainerLeft = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'row',
    gap: '8px',
    alignItems: 'center',
  },
})

interface StatusCheckboxProps {
  apps: Application[] | undefined
  state: ApplicationState
  title: string
}

export const StatusCheckbox = (props: StatusCheckboxProps): JSX.Element => {
  const num = () => props.apps?.filter((app) => applicationState(app) === props.state)?.length ?? 0
  return (
    <Container>
      <ContainerLeft>
        <StatusIcon state={props.state} />
        <div>{props.title}</div>
      </ContainerLeft>
      <div>{num()}</div>
    </Container>
  )
}
