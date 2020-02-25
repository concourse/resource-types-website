module ResourceList.View exposing (view)

import Card.View as CardView exposing (view)
import Common.Common exposing (ResourceType)
import Common.Overrides as Overrides exposing (grid)
import Element exposing (Element, fill, height, maximum, paddingXY, width, wrappedRow)
import List exposing (sortBy)
import ResourceList.ResourceList exposing (container)


view : List ResourceType -> String -> String -> Element msg
view resourceList githubIconImg githubStarImg =
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
        (List.map (viewCard githubIconImg githubStarImg) sortedRL)


viewCard : String -> String -> ResourceType -> Element msg
viewCard githubIconImg githubStarImg resourceType =
    CardView.view resourceType githubIconImg githubStarImg
