module Footer.Footer exposing (container, footer, link)

import Common.Common as Common
import Footer.Styles as Styles


type alias Footer =
    { container : Container
    , link : Link
    }


type alias Container =
    { height : Int
    , backgroundColor : Common.RGB
    , spacing : Int
    }


type alias Link =
    { font : String
    , size : Int
    , color : Common.RGB
    }


footer : Footer
footer =
    { container = container, link = link }


container : Container
container =
    { height = Styles.footerHeight
    , backgroundColor = Styles.footerBackgroundColor
    , spacing = Styles.footerContentSpacing
    }


link : Link
link =
    { font = Styles.linkFont
    , size = Styles.linkSize
    , color = Styles.linkColor
    }
