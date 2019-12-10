module Common.Overrides exposing (ellipsis, grid)

import Element as Element exposing (htmlAttribute)
import Html
import Html.Attributes exposing (style)


grid : List (Element.Attribute msg)
grid =
    [ htmlAttribute (Html.Attributes.style "display" "grid")
    , htmlAttribute (Html.Attributes.style "grid-gap" "16px")
    , htmlAttribute (Html.Attributes.style "justify-content" "center")
    , htmlAttribute (Html.Attributes.style "grid-template-columns" "repeat(auto-fill, 280px)")
    , htmlAttribute (Html.Attributes.style "margin" "75px auto")
    ]


ellipsis : List (Html.Attribute msg)
ellipsis =
    [ Html.Attributes.style "text-overflow" "ellipsis"
    , Html.Attributes.style "white-space" "nowrap"
    , Html.Attributes.style "overflow" "hidden"
    , Html.Attributes.style "height" "26px"
    ]
