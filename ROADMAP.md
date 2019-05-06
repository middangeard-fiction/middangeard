Middangeard Roadmap
===========

## v1

* Basic adventure creation toolkit
* Focus on simplicity
* Refined audio system (Operatmos)

## v2

* Replace strings (such as the ones used for room hashes) 
  with full fledged "objects"
* Code clean-up
* Extended ``Player`` attributes

## v3

* Lua VM for advanced scripting
* "Narrator Mode". If enabled, will load a game's **narration.ogg**. Authors will record the full
  narration into this one file, then play back specific parts through **game.Narrate(fromTime, toTime)**