extends Node


# Called when the node enters the scene tree for the first time.
func _ready():
	var myMesh = MeshInstance3D.new()
	myMesh.set_mesh(BoxMesh.new())
	add_child(myMesh)


# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta):
	pass
