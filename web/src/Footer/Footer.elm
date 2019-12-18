module Footer.Footer exposing (..)

import Common.Common as Common
import Footer.Styles as Styles


type alias Footer =
    { container : Container
    }


type alias Container =
    { height : Int
    , backgroundColor : Common.RGB
    , font : String
    , size : Int
    , color : Common.RGB
    }


footer : Footer
footer =
    { container = container }


container : Container
container =
    { height = Styles.footerHeight
    , backgroundColor = Styles.footerBackgroundColor
    , font = Styles.footerFont
    , size = Styles.footerSize
    , color = Styles.footerColor
    }
