module Footer.Styles exposing (footerBackgroundColor, footerColor, footerFont, footerHeight, footerSize)

import Common.Common as Common


footerHeight : Int
footerHeight =
    Common.gridSize * 5


footerBackgroundColor : Common.RGB
footerBackgroundColor =
    Common.footerBackgroundColor


footerFont : String
footerFont =
    "Barlow"


footerSize : Int
footerSize =
    14


footerColor : Common.RGB
footerColor =
    Common.white
