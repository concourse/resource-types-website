module Main exposing (view)

import Banner.View as Banner exposing (view)
import Html exposing (..)
import Html.Attributes exposing (class, href, rel, style)


main =
    div [] [ view ]


view : Html msg
view =
    div
        [ class "wrapper"
        , style "margin" "0 auto"
        , style "font-face" "Roboto Slab"
        ]
        [ Banner.view
        ]
