module Banner.Styles exposing
    ( backgroundColor
    , bannerHeight
    , bodyColor
    , bodyFont
    , bodySize
    , bodyWidth
    , titleColor
    , titleFont
    , titleLineHeight
    , titlePaddingBottom
    , titlePaddingTop
    , titleSize
    )

import Common.Common as Common



-- CONTAINER


bannerHeight : Int
bannerHeight =
    Common.gridSize * 24


backgroundColor : Common.RGB
backgroundColor =
    Common.bannerBackgroundColor



-- TITLE


titleFont : String
titleFont =
    "Roboto Slab"


titleColor : Common.RGB
titleColor =
    Common.white


titleSize : Int
titleSize =
    26


titleLineHeight : Int
titleLineHeight =
    42


titlePaddingTop : Int
titlePaddingTop =
    Common.gridSize * 8


titlePaddingBottom : Int
titlePaddingBottom =
    Common.gridSize * 2



-- BODY


bodyFont : String
bodyFont =
    "Barlow"


bodyColor : Common.RGB
bodyColor =
    Common.white


bodySize : Int
bodySize =
    18


bodyWidth : Int
bodyWidth =
    Common.gridSize * 50
