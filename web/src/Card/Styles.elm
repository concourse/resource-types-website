module Card.Styles exposing (borderRadius, height, hoverShadow, shadow, spacing, width)

import Common.Common as Common exposing (Shadow)


height : Int
height =
    160


width : Int
width =
    280


borderRadius : Int
borderRadius =
    4


spacing : Int
spacing =
    16


shadow : Shadow
shadow =
    { offsetX = 0
    , offsetY = 2
    , blur = 3
    , size = 1
    , color = Common.greyishRed
    }


hoverShadow : Shadow
hoverShadow =
    { offsetX = 1
    , offsetY = 3
    , blur = 8
    , size = 1
    , color = Common.greyishRed
    }
