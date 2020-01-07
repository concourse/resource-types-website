module Terms.Terms exposing (Terms, body, container, terms, title)

import Terms.Styles as Styles


type alias Terms =
    { container : Container
    , title : Title
    , body : Body
    }


type alias Container =
    { width : Int
    , spacing : Int
    }


type alias Title =
    { font : String
    , size : Int
    }


type alias Body =
    { font : String
    , size : Int
    }


terms : Terms
terms =
    { container = container, title = title, body = body }


container : Container
container =
    { width = Styles.containerWidth, spacing = Styles.containerSpacing }


title : Title
title =
    { size = Styles.titleSize, font = Styles.titleFont }


body : Body
body =
    { size = Styles.bodySize, font = Styles.bodyFont }
