components {
  id: "note"
  component: "/main/note.script"
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
  position {
    z: 1.0
  }
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
    z: 1.0
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
    z: 1.0
  }
  scale {
    x: 0.5
    y: 0.5
  }
}
embedded_components {
  id: "collisionobject3"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"3\"\n"
  "mask: \"trigger\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 160.0\n"
  "      y: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"collider\"\n"
  "  }\n"
  "  data: 34.086895\n"
  "  data: 33.074753\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject1"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"1\"\n"
  "mask: \"trigger\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"collider\"\n"
  "  }\n"
  "  data: 34.086895\n"
  "  data: 33.074753\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject2"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"2\"\n"
  "mask: \"trigger\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 80.0\n"
  "      y: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"collider\"\n"
  "  }\n"
  "  data: 34.086895\n"
  "  data: 33.074753\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
