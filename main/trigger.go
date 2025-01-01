components {
  id: "trigger"
  component: "/main/trigger.script"
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"logo\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 192.0\n"
  "  y: 192.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.5
    y: 0.5
  }
}
embedded_components {
  id: "sprite2"
  type: "sprite"
  data: "default_animation: \"logo\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 192.0\n"
  "  y: 192.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  ""
  position {
    x: 80.0
  }
  scale {
    x: 0.5
    y: 0.5
  }
}
embedded_components {
  id: "sprite3"
  type: "sprite"
  data: "default_animation: \"logo\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 192.0\n"
  "  y: 192.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  ""
  position {
    x: 160.0
  }
  scale {
    x: 0.5
    y: 0.5
  }
}
embedded_components {
  id: "collisionobject3"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"trigger\"\n"
  "mask: \"3\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 160.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"3\"\n"
  "  }\n"
  "  data: 10.0\n"
  "  data: 25.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject1"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"trigger\"\n"
  "mask: \"1\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"1\"\n"
  "  }\n"
  "  data: 10.0\n"
  "  data: 25.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject2"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"trigger\"\n"
  "mask: \"2\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 80.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"2\"\n"
  "  }\n"
  "  data: 10.0\n"
  "  data: 25.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
