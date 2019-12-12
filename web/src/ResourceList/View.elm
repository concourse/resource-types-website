module ResourceList.View exposing (view)

import Card.View as CardView exposing (view)
import Common.Common exposing (ResourceType)
import Common.Overrides as Overrides exposing (grid)
import Element exposing (Element, fill, maximum, paddingXY, width, wrappedRow)
import ResourceList.ResourceList exposing (container)

view : List ResourceType -> Element msg
view resourceList =
    wrappedRow
        ([ width (fill |> maximum container.maxWidth)
         , paddingXY container.outsideMargin 0
         ]
            ++ Overrides.grid
        )
        (List.map viewCard resourceList)


viewCard resourceType = 
    CardView.view resourceType
