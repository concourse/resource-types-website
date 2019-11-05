module Main exposing (view)

import Banner.Banner as Banner exposing (view)
import Html exposing (..)
import Html.Attributes exposing (class, href, rel, style)


main =
    div []
        [ view
        ]


view : Html msg
view =
    div
        [ class "wrapper"
        , style "width" "1024px"
        , style "margin" "0 auto"
        ]
        [ node "link"
            [ href "https://fonts.googleapis.com/css?family=Jura", rel "stylesheet" ]
            []
        , Banner.view
        ]
