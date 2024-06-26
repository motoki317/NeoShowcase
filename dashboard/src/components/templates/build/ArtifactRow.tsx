import { styled } from '@macaron-css/solid'
import type { Component } from 'solid-js'
import type { Artifact } from '/@/api/neoshowcase/protobuf/gateway_pb'
import { Button } from '/@/components/UI/Button'
import { client, handleAPIError } from '/@/libs/api'
import { saveToFile } from '/@/libs/download'
import { formatBytes } from '/@/libs/format'
import { colorVars, textVars } from '/@/theme'

const Container = styled('div', {
  base: {
    width: '100%',
    padding: '16px 16px 16px 20px',
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    gap: '8px',
    background: colorVars.semantic.ui.primary,
  },
})
const ContentsContainer = styled('div', {
  base: {
    width: '100%',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'flex-start',
  },
})
const TitleContainer = styled('div', {
  base: {
    width: '100%',
    display: 'flex',
    alignItems: 'center',
    gap: '8px',
  },
})
const ArtifactName = styled('div', {
  base: {
    width: '100%',
    overflow: 'hidden',
    textOverflow: 'ellipsis',
    whiteSpace: 'nowrap',
    color: colorVars.semantic.text.black,
    ...textVars.h4.regular,
  },
})
const MetaContainer = styled('div', {
  base: {
    width: '100%',
    display: 'flex',
    alignItems: 'center',
    gap: '4px',

    color: colorVars.semantic.text.grey,
    ...textVars.caption.regular,
  },
})
const ArtifactSize = styled('div', {
  base: {
    width: 'fit-content',
    overflow: 'hidden',
    textOverflow: 'ellipsis',
    whiteSpace: 'nowrap',
  },
})
export interface Props {
  artifact: Artifact
}

const downloadArtifact = async (id: string) => {
  try {
    const data = await client.getBuildArtifact({ artifactId: id })
    saveToFile(data.content, 'application/gzip', data.filename)
  } catch (e) {
    handleAPIError(e, '成果物のダウンロードに失敗しました')
  }
}

export const ArtifactRow: Component<Props> = (props) => {
  return (
    <Container>
      <ContentsContainer>
        <TitleContainer>
          <ArtifactName>{props.artifact.name}</ArtifactName>
        </TitleContainer>
        <MetaContainer>
          <ArtifactSize>{formatBytes(+props.artifact.size.toString())}</ArtifactSize>
        </MetaContainer>
      </ContentsContainer>
      <Button variants="ghost" size="medium" onClick={() => downloadArtifact(props.artifact.id)}>
        Download
      </Button>
    </Container>
  )
}
