module ResourceList.View exposing (view)

import Card.View as CardView exposing (view)
import Common.Common exposing (ResourceType)
import Common.Overrides as Overrides exposing (grid)
import Element exposing (Element, fill, height, maximum, paddingXY, width, wrappedRow)
import List exposing (sortBy)
import ResourceList.ResourceList exposing (container)


view : List ResourceType -> String -> Element msg
view resourceList githubIconImg =
    let
        sortedRL =
            sortBy .name resourceList
    in
    wrappedRow
        ([ width (fill |> maximum container.maxWidth)
         , height fill
         , paddingXY container.outsideMargin 0
         ]
            ++ Overrides.grid
        )
        (List.map (viewCard githubIconImg) sortedRL)


viewCard : String -> ResourceType -> Element msg
viewCard flags resourceType =
    CardView.view resourceType flags
