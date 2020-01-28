module Terms.Terms exposing (Terms, backLink, body, containerWidth, terms, title)

import Common.Common as Common
import Terms.Styles as Styles


type alias Terms =
    { containerWidth : Int
    , backLink : BackLink
    , title : Title
    , body : Body
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
    { containerWidth = containerWidth
    , backLink = backLink
    , title = title
    , body = body
    }


containerWidth : Int
containerWidth =
    Styles.containerWidth


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
