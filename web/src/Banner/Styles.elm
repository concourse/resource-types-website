module Banner.Styles exposing
    ( backgroundColor
    , backgroundImage
    , bannerHeight
    , bodyColor
    , bodyFont
    , bodySize
    , bodyWidth
    , titleColor
    , titleFont
    , titleLineHeight
    , titleSize
    )

import Common.Common as Common



-- CONTAINER


bannerHeight : Int
bannerHeight =
    176


backgroundImage : String
backgroundImage =
    "banner-background.png"


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
    24


titleLineHeight : Int
titleLineHeight =
    32



-- BODY


bodyFont : String
bodyFont =
    "Barlow"


bodyColor : Common.RGB
bodyColor =
    Common.white


bodySize : Int
bodySize =
    16


bodyWidth : Int
bodyWidth =
    400
