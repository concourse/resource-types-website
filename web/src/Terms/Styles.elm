module Terms.Styles exposing
    ( bodyFont
    , bodySize
    , containerSpacing
    , containerWidth
    , titleFont
    , titleSize
    )

import Common.Common as Common



-- Container


containerWidth : Int
containerWidth =
    Common.gridSize * 65


containerSpacing : Int
containerSpacing =
    Common.gridSize * 9



-- Title


titleSize : Int
titleSize =
    32


titleFont : String
titleFont =
    "Roboto Slab"



-- Body


bodySize : Int
bodySize =
    14


bodyFont : String
bodyFont =
    "Barlow"
