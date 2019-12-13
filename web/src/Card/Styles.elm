module Card.Styles exposing
    ( containerBorderRadius
    , containerHeight
    , containerHoverShadow
    , containerPaddingLeft
    , containerShadow
    , containerSpacing
    , containerWidth
    , descriptionColor
    , descriptionFont
    , descriptionMaxHeight
    , descriptionMaxWidth
    , descriptionMinHeight
    , descriptionPaddingTop
    , descriptionSize
    , descriptionSpacing
    , githubImageHeight
    , githubImageName
    , githubImagePaddingTop
    , githubImageWidth
    , nameColor
    , nameFont
    , nameMaxWidth
    , namePaddingTop
    , nameSize
    )

import Common.Common as Common exposing (RGB, Shadow)



-- Container


containerHeight : Int
containerHeight =
    Common.gridSize * 20


containerWidth : Int
containerWidth =
    Common.gridSize * 35


containerBorderRadius : Int
containerBorderRadius =
    4


containerSpacing : Int
containerSpacing =
    Common.gridSize * 2


containerPaddingLeft : Int
containerPaddingLeft =
    Common.gridSize * 2


containerShadow : Shadow
containerShadow =
    { offsetX = 0
    , offsetY = 2
    , blur = 3
    , size = 1
    , color = Common.shadowColor
    }


containerHoverShadow : Shadow
containerHoverShadow =
    { offsetX = 1
    , offsetY = 3
    , blur = 8
    , size = 1
    , color = Common.shadowColor
    }



-- resource type
-- name


nameSize : Int
nameSize =
    20


nameFont : String
nameFont =
    "Roboto Slab"


nameColor : RGB
nameColor =
    Common.cardTitleColor


namePaddingTop : Int
namePaddingTop =
    Common.gridSize * 3


nameMaxWidth : Int
nameMaxWidth =
    containerWidth - (containerSpacing * 2)



-- description


descriptionSize : Int
descriptionSize =
    12


descriptionFont : String
descriptionFont =
    "Barlow"


descriptionPaddingTop : Int
descriptionPaddingTop =
    Common.gridSize * 3


descriptionColor : RGB
descriptionColor =
    Common.cardDescriptionColor


descriptionMaxWidth : Int
descriptionMaxWidth =
    containerWidth - (containerSpacing * 2)


descriptionSpacing : Int
descriptionSpacing =
    descriptionSize - Common.gridSize



-- the height of the two lines of text and the spacing above, between, and below each line


descriptionMaxHeight : Int
descriptionMaxHeight =
    Common.gridSize * 7


descriptionMinHeight : Int
descriptionMinHeight =
    Common.gridSize * 7



-- github


githubImageName : String
githubImageName =
    "github-logo.png"


githubImageWidth : Int
githubImageWidth =
    Common.gridSize * 2


githubImageHeight : Int
githubImageHeight =
    Common.gridSize * 2


githubImagePaddingTop : Int
githubImagePaddingTop =
    Common.gridSize * 3
