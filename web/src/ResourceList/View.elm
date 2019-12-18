module ResourceList.View exposing (view)

import Card.View as CardView exposing (view)
import Common.Common exposing (ResourceType)
import Common.Overrides as Overrides exposing (grid)
import Element exposing (Element, fill, maximum, paddingXY, width, wrappedRow)
import ResourceList.ResourceList exposing (container)


view : List ResourceType -> String -> Element msg
view resourceList githubIconImg =
    wrappedRow
        ([ width (fill |> maximum container.maxWidth)
         , paddingXY container.outsideMargin 0
         ]
            ++ Overrides.grid
        )
        (List.map (viewCard githubIconImg) resourceList)


viewCard : String -> ResourceType -> Element msg
viewCard flags resourceType =
    CardView.view resourceType flags
