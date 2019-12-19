module Footer.Styles exposing
    ( footerBackgroundColor
    , footerContentSpacing
    , footerHeight
    , linkColor
    , linkFont
    , linkSize
    )

import Common.Common as Common


footerHeight : Int
footerHeight =
    Common.gridSize * 5


footerBackgroundColor : Common.RGB
footerBackgroundColor =
    Common.footerBackgroundColor


footerContentSpacing : Int
footerContentSpacing =
    12


linkFont : String
linkFont =
    "Barlow"


linkSize : Int
linkSize =
    14


linkColor : Common.RGB
linkColor =
    Common.white
