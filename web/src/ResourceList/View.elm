module ResourceList.View exposing (view)

import Card.View as CardView exposing (view)
import Element exposing (Element, fill, htmlAttribute, maximum, width, wrappedRow)
import Html.Attributes exposing (style)
import ResourceList.ResourceList exposing (container)


view : Element msg
view =
    wrappedRow
        [ width (fill |> maximum container.maxWidth)

        -- TODO figure out how to do this in elm-ui
        , htmlAttribute (Html.Attributes.style "display" "grid")
        , htmlAttribute (Html.Attributes.style "grid-gap" "16px")
        , htmlAttribute (Html.Attributes.style "justify-content" "center")
        , htmlAttribute (Html.Attributes.style "grid-template-columns" "repeat(auto-fill, 280px)")
        , htmlAttribute (Html.Attributes.style "margin" "75px auto")
        ]
        [ CardView.view, CardView.view, CardView.view, CardView.view, CardView.view, CardView.view ]
