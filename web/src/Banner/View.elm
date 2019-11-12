module Banner.View exposing (view)

import Banner.Styles as Styles exposing (body, container, title)
import Html exposing (Html, div, text)
import Html.Attributes exposing (class)


view : Html msg
view =
    div
        (class "banner-container" :: Styles.container)
        [ div
            (class "banner-title" :: Styles.title)
            [ Html.text "Concourse Resources" ]
        , div
            (class "banner-body" :: Styles.body)
            [ Html.text "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt"
            ]
        ]
