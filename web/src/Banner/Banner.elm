module Banner.Banner exposing (Banner, Body, Container, Title, banner, body, container, title)

import Banner.Styles as Styles
    exposing
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
import Common.Common exposing (RGB)


type alias Banner =
    { container : Container
    , title : Title
    , body : Body
    }


type alias Container =
    { height : Int
    , backgroundColor : RGB
    , backgroundImage : String
    }


type alias Title =
    { size : Int
    , color : RGB
    , text : String
    , font : String
    , lineHeight : Int
    }


type alias Body =
    { size : Int
    , color : RGB
    , text : String
    , font : String
    , width : Int
    }


banner : Banner
banner =
    { container = container
    , title = title
    , body = body
    }


container : Container
container =
    { height = Styles.bannerHeight
    , backgroundColor = Styles.backgroundColor
    , backgroundImage = Styles.backgroundImage
    }


title : Title
title =
    { size = Styles.titleSize
    , color = Styles.titleColor
    , text = "Concourse Resources"
    , font = Styles.titleFont
    , lineHeight = Styles.titleLineHeight
    }


body : Body
body =
    { size = Styles.bodySize
    , color = Styles.bodyColor
    , text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt"
    , font = Styles.bodyFont
    , width = Styles.bodyWidth
    }
