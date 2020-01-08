module Terms.Terms exposing (Terms, backLink, body, container, terms, title)

import Common.Common as Common
import Terms.Styles as Styles


type alias Terms =
    { container : Container
    , backLink : BackLink
    , title : Title
    , body : Body
    }


type alias Container =
    { width : Int
    }


type alias BackLink =
    { size : Int
    , paddingTop : Int
    , color : Common.RGB
    }


type alias Title =
    { font : String
    , size : Int
    , padding : Int
    }


type alias Body =
    { font : String
    , size : Int
    }


terms : Terms
terms =
    { container = container
    , backLink = backLink
    , title = title
    , body = body
    }


container : Container
container =
    { width = Styles.containerWidth }


backLink : BackLink
backLink =
    { size = Styles.backLinkSize
    , paddingTop = Styles.backLinkPaddingTop
    , color = Styles.backLinkColor
    }


title : Title
title =
    { size = Styles.titleSize, font = Styles.titleFont, padding = Styles.titlePadding }


body : Body
body =
    { size = Styles.bodySize, font = Styles.bodyFont }
