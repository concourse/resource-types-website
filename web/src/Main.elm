module Main exposing (view)

import Html exposing (..)
import Html.Attributes exposing (..)


main =
    view model


model =
    {}


view : {} -> Html msg
view viewModel =
    div []
        [ div
            [ class "banner-title"
            , style "height" "176px"
            , style "text-align" "center"
            ]
            [ Html.text "Concourse Resources" ]
        , div
            [ class "banner-body" ]
            [ Html.text "" ]
        ]
