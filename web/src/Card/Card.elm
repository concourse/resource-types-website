module Card.Card exposing (Card, Container, Description, Github, Name, ResourceType, card, container, resourceType)

import Card.Styles as Styles
import Common.Common exposing (RGB, Shadow)


type alias Card =
    { container : Container
    , resourceType : ResourceType
    }


type alias Container =
    { height : Int
    , width : Int
    , borderRadius : Int
    , shadow : Shadow
    , hoverShadow : Shadow
    , spacing : Int
    , paddingLeft : Int
    }


type alias ResourceType =
    { name : Name
    , description : Description
    , github : Github
    }


type alias Name =
    { size : Int
    , font : String
    , paddingTop : Int
    , maxWidth : Int
    , color : RGB
    }


type alias Description =
    { size : Int
    , font : String
    , paddingTop : Int
    , color : RGB
    , maxHeight : Int
    , minHeight : Int
    , maxWidth : Int
    , spacing : Int
    }


type alias Github =
    { imageName : String
    , imageHeight : Int
    , imageWidth : Int
    , paddingTop : Int
    }


card : Card
card =
    { container = container
    , resourceType = resourceType
    }


container : Container
container =
    { height = Styles.containerHeight
    , width = Styles.containerWidth
    , borderRadius = Styles.containerBorderRadius
    , shadow = Styles.containerShadow
    , hoverShadow = Styles.containerHoverShadow
    , spacing = Styles.containerSpacing
    , paddingLeft = Styles.containerPaddingLeft
    }


resourceType : ResourceType
resourceType =
    { name = name
    , description = description
    , github = github
    }


name : Name
name =
    { size = Styles.nameSize
    , font = Styles.nameFont
    , paddingTop = Styles.namePaddingTop
    , maxWidth = Styles.nameMaxWidth
    , color = Styles.nameColor
    }


description : Description
description =
    { size = Styles.descriptionSize
    , font = Styles.descriptionFont
    , paddingTop = Styles.descriptionPaddingTop
    , color = Styles.descriptionColor
    , maxHeight = Styles.descriptionMaxHeight
    , minHeight = Styles.descriptionMinHeight
    , maxWidth = Styles.descriptionMaxWidth
    , spacing = Styles.descriptionSpacing
    }


github : Github
github =
    { imageName = Styles.githubImageName
    , imageHeight = Styles.githubImageHeight
    , imageWidth = Styles.githubImageWidth
    , paddingTop = Styles.githubImagePaddingTop
    }
