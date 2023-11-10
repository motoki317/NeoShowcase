import { Application, Repository, UpdateRepositoryRequest } from '/@/api/neoshowcase/protobuf/gateway_pb'
import { Button } from '/@/components/UI/Button'
import { TextInput } from '/@/components/UI/TextInput'
import { DataTable } from '/@/components/layouts/DataTable'
import FormBox from '/@/components/layouts/FormBox'
import { FormItem } from '/@/components/templates/FormItem'
import { client, handleAPIError } from '/@/libs/api'
import { providerToIcon, repositoryURLToProvider } from '/@/libs/application'
import useModal from '/@/libs/useModal'
import { useRepositoryData } from '/@/routes'
import { colorVars, textVars } from '/@/theme'
import { PlainMessage } from '@bufbuild/protobuf'
import { styled } from '@macaron-css/solid'
import { SubmitHandler, createForm, required, reset } from '@modular-forms/solid'
import { useNavigate } from '@solidjs/router'
import { Component, Show, createEffect, on } from 'solid-js'
import toast from 'solid-toast'

type GeneralForm = Required<Pick<PlainMessage<UpdateRepositoryRequest>, 'name'>>

const NameConfig: Component<{
  repo: Repository
  refetchRepo: () => void
}> = (props) => {
  const [generalForm, General] = createForm<GeneralForm>({
    initialValues: {
      name: props.repo.name,
    },
  })

  createEffect(() => {
    reset(generalForm, 'name', {
      initialValue: props.repo.name,
    })
  })

  const handleSubmit: SubmitHandler<GeneralForm> = async (values) => {
    try {
      await client.updateRepository({
        id: props.repo.id,
        name: values.name,
      })
      toast.success('Project名を更新しました')
      props.refetchRepo()
    } catch (e) {
      handleAPIError(e, 'Project名の更新に失敗しました')
    }
  }
  const discardChanges = () => {
    reset(generalForm)
  }

  return (
    <General.Form onSubmit={handleSubmit}>
      <FormBox.Container>
        <FormBox.Forms>
          <General.Field name="name" validate={[required('Enter Project Name')]}>
            {(field, props) => (
              <FormItem title="Project Name" required>
                <TextInput value={field.value} error={field.error} {...props} />
              </FormItem>
            )}
          </General.Field>
        </FormBox.Forms>
        <FormBox.Actions>
          <Show when={generalForm.dirty && !generalForm.submitting}>
            <Button variants="borderError" size="small" onClick={discardChanges} type="button">
              Discard Changes
            </Button>
          </Show>
          <Button
            variants="primary"
            size="small"
            type="submit"
            disabled={generalForm.invalid || !generalForm.dirty || generalForm.submitting}
          >
            Save
          </Button>
        </FormBox.Actions>
      </FormBox.Container>
    </General.Form>
  )
}

const DeleteProjectNotice = styled('div', {
  base: {
    color: colorVars.semantic.text.grey,
    ...textVars.caption.regular,
  },
})
const DeleteConfirm = styled('div', {
  base: {
    width: '100%',
    padding: '16px 20px',
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    gap: '8px',
    borderRadius: '8px',
    background: colorVars.semantic.ui.secondary,
    color: colorVars.semantic.text.black,
    ...textVars.h3.regular,
  },
})

const DeleteProject: Component<{
  repo: Repository
  apps: Application[]
}> = (props) => {
  const { Modal, open, close } = useModal()
  const navigate = useNavigate()

  const deleteRepository = async () => {
    try {
      await client.deleteRepository({ repositoryId: props.repo.id })
      toast.success('Projectを削除しました')
      close()
      navigate('/apps')
    } catch (e) {
      handleAPIError(e, 'Projectの削除に失敗しました')
    }
  }
  const canDeleteRepository = () => props.apps.length === 0

  return (
    <>
      <FormBox.Container>
        <FormBox.Forms>
          <FormItem title="Delete Project">
            <DeleteProjectNotice>
              Projectを削除するには、このプロジェクト内のすべてのAppを削除する必要があります。
            </DeleteProjectNotice>
          </FormItem>
        </FormBox.Forms>
        <FormBox.Actions>
          <Button
            variants="primaryError"
            size="small"
            onClick={open}
            type="button"
            disabled={!canDeleteRepository()}
            tooltip={{
              props: {
                content: !canDeleteRepository() ? 'Project内にAppが存在するため削除できません' : undefined,
              },
            }}
          >
            Delete Project
          </Button>
        </FormBox.Actions>
      </FormBox.Container>
      <Modal.Container>
        <Modal.Header>Delete Repository</Modal.Header>
        <Modal.Body>
          <DeleteConfirm>
            {providerToIcon(repositoryURLToProvider(props.repo.url), 24)}
            {props.repo.name}
          </DeleteConfirm>
        </Modal.Body>
        <Modal.Footer>
          <Button variants="text" size="medium" onClick={close} type="button">
            No, Cancel
          </Button>
          <Button variants="primaryError" size="medium" onClick={deleteRepository} type="button">
            Yes, Delete
          </Button>
        </Modal.Footer>
      </Modal.Container>
    </>
  )
}

export default () => {
  const { repo, refetchRepo, apps } = useRepositoryData()
  const loaded = () => !!(repo() && apps())

  return (
    <DataTable.Container>
      <DataTable.Title>General</DataTable.Title>
      <Show when={loaded()}>
        <NameConfig repo={repo()} refetchRepo={refetchRepo} />
        <DeleteProject repo={repo()} apps={apps()} />
      </Show>
    </DataTable.Container>
  )
}
