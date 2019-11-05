module Main exposing (main, view)

import Banner.View as Banner exposing (view)
import Html exposing (Html, div)
import Html.Attributes exposing (class, style)

main : Html msg
main =
    div [] [ view ]


view : Html msg
view =
    div
        [ class "wrapper"
        , style "margin" "0 auto"
        ]
        [ Banner.view
        ]
