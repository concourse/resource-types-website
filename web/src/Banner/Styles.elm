module Banner.Styles exposing
    ( body
    , container
    , title
    )

import Html
import Html.Attributes exposing (style)


container : List (Html.Attribute msg)
container =
    [ style "height" "176px"
    , style "display" "grid"
    , style "background-color" "#2A3239"
    , style "background-image" "url(banner-background.png)"
    , style "background-size" "100%"
    ]


title : List (Html.Attribute msg)
title =
    [ style "font-size" "24px"
    , style "display" "grid"
    , style "align-items" "center"
    , style "justify-content" "center"
    , style "font-family" "Roboto Slab"
    , style "color" "#FFFFFF"
    , style "line-height" "32px"
    ]


body : List (Html.Attribute msg)
body =
    [ style "display" "grid"
    , style "justify-content" "center"
    , style "font-family" "Barlow"
    , style "color" "#FFFFFF"
    , style "grid-template-columns" "minmax(auto, 400px)"
    , style "text-align" "center"
    , style "line-height" "24px"
    ]
