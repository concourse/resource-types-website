module Card.Styles exposing
    ( authorColor
    , authorFont
    , authorPaddingTop
    , authorSize
    , containerBorderRadius
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
    , hostImageHeight
    , hostImagePaddingTop
    , hostImageWidth
    , hostPillBorderRadius
    , hostPillDarkBackgroundColor
    , hostPillFont
    , hostPillFontSize
    , hostPillHeight
    , hostPillImageHeight
    , hostPillImageWidth
    , hostPillLightBackgroundColor
    , hostPillPaddingLeft
    , hostPillPaddingRight
    , hostPillSpacing
    , hostSpacing
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
    { offset = ( 0, 2 )
    , blur = 3
    , size = 1
    , color = Common.shadowColor
    }


containerHoverShadow : Shadow
containerHoverShadow =
    { offset = ( 1, 3 )
    , blur = 8
    , size = 1
    , color = Common.shadowColor
    }



-- resource type
-- name


nameSize : Int
nameSize =
    16


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



-- author


authorFont : String
authorFont =
    "Roboto Slab"


authorSize : Int
authorSize =
    12


authorColor : RGB
authorColor =
    Common.cardTitleColor


authorPaddingTop : Int
authorPaddingTop =
    Common.gridSize // 4



-- description


descriptionSize : Int
descriptionSize =
    14


descriptionFont : String
descriptionFont =
    "Barlow"


descriptionPaddingTop : Int
descriptionPaddingTop =
    Common.gridSize


descriptionColor : RGB
descriptionColor =
    Common.cardDescriptionColor


descriptionMaxWidth : Int
descriptionMaxWidth =
    containerWidth - (containerSpacing * 2)


descriptionSpacing : Int
descriptionSpacing =
    Common.gridSize // 4


descriptionMaxHeight : Int
descriptionMaxHeight =
    Common.gridSize * 4 + descriptionSize * 2


descriptionMinHeight : Int
descriptionMinHeight =
    descriptionMaxHeight



-- host


hostSpacing : Int
hostSpacing =
    3



--image


hostImageWidth : Int
hostImageWidth =
    Common.gridSize * 2


hostImageHeight : Int
hostImageHeight =
    Common.gridSize * 2


hostImagePaddingTop : Int
hostImagePaddingTop =
    Common.gridSize



-- pill


hostPillLightBackgroundColor : RGB
hostPillLightBackgroundColor =
    Common.hostPillLightBackgroundColor


hostPillDarkBackgroundColor : RGB
hostPillDarkBackgroundColor =
    Common.hostPillDarkBackgroundColor


hostPillHeight : Int
hostPillHeight =
    Common.gridSize * 2


hostPillFontSize : Int
hostPillFontSize =
    10


hostPillFont : String
hostPillFont =
    "iosevka"


hostPillPaddingLeft : Int
hostPillPaddingLeft =
    Common.gridSize // 2


hostPillPaddingRight : Int
hostPillPaddingRight =
    (Common.gridSize // 2) + 2


hostPillSpacing : Int
hostPillSpacing =
    3


hostPillBorderRadius : Int
hostPillBorderRadius =
    5


hostPillImageHeight : Int
hostPillImageHeight =
    10


hostPillImageWidth : Int
hostPillImageWidth =
    10
