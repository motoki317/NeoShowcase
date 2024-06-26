import { Show } from 'solid-js'
import toast from 'solid-toast'
import type { User } from '/@/api/neoshowcase/protobuf/gateway_pb'
import { DataTable } from '/@/components/layouts/DataTable'
import OwnerList from '/@/components/templates/OwnerList'
import { client, handleAPIError } from '/@/libs/api'
import { userFromId, users } from '/@/libs/useAllUsers'
import { useApplicationData } from '/@/routes'

export default () => {
  const { app, refetch, hasPermission } = useApplicationData()
  const loaded = () => !!(app() && users())

  const handleAddOwner = async (user: User) => {
    const newOwnerIds = app()?.ownerIds.concat(user.id)
    try {
      await client.updateApplication({
        id: app()?.id,
        ownerIds: { ownerIds: newOwnerIds },
      })
      toast.success('アプリケーションオーナーを追加しました')
      void refetch()
    } catch (e) {
      handleAPIError(e, 'アプリケーションオーナーの追加に失敗しました')
    }
  }

  const handleDeleteOwner = async (user: User) => {
    const newOwnerIds = app()?.ownerIds.filter((id) => id !== user.id)
    try {
      await client.updateApplication({
        id: app()?.id,
        ownerIds: { ownerIds: newOwnerIds },
      })
      toast.success('アプリケーションオーナーを削除しました')
      void refetch()
    } catch (e) {
      handleAPIError(e, 'アプリケーションオーナーの削除に失敗しました')
    }
  }

  return (
    <DataTable.Container>
      <DataTable.Title>Owners</DataTable.Title>
      <DataTable.SubTitle>
        オーナーはアプリ設定の変更, アプリログ/メトリクスの閲覧, 環境変数の閲覧, ビルドログの閲覧が可能になります
      </DataTable.SubTitle>
      <Show when={loaded()}>
        <OwnerList
          owners={app()!.ownerIds.map(userFromId)}
          users={users()!}
          handleAddOwner={handleAddOwner}
          handleDeleteOwner={handleDeleteOwner}
          hasPermission={hasPermission()}
        />
      </Show>
    </DataTable.Container>
  )
}
