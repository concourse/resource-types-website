module Main exposing (banner, layout, main, resourceList)

import Banner.View as Banner exposing (view)
import Element exposing (Element, column, fill, width)
import Html exposing (Html)
import ResourceList.View as ResourceList exposing (view)


main : Html msg
main =
    Element.layout [] layout


banner : Element msg
banner =
    Banner.view


resourceList : Element msg
resourceList =
    ResourceList.view


layout : Element msg
layout =
    column
        [ width fill ]
        [ banner
        , resourceList
        ]
