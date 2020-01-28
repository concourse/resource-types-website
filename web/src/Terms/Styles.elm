module Terms.Styles exposing
    ( backLinkColor
    , backLinkPaddingTop
    , backLinkSize
    , bodyFont
    , bodySize
    , containerWidth
    , titleFont
    , titlePadding
    , titleSize
    )

import Common.Common as Common



-- Container


containerWidth : Int
containerWidth =
    Common.gridSize * 65



-- Back Link


backLinkSize : Int
backLinkSize =
    14


backLinkPaddingTop : Int
backLinkPaddingTop =
    Common.gridSize * 6


backLinkColor : Common.RGB
backLinkColor =
    Common.termsLinkColor



-- Title


titleSize : Int
titleSize =
    32


titleFont : String
titleFont =
    "Roboto Slab"


titlePadding : Int
titlePadding =
    Common.gridSize * 3



-- Body


bodySize : Int
bodySize =
    14


bodyFont : String
bodyFont =
    "Barlow"
