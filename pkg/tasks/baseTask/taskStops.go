package dtasksbase

type DTaskStop uint

const MatchmakerStop DTaskStop = iota
const ListenerStop DTaskStop = iota
const BroadcasterStop DTaskStop = iota
const ButlerStop DTaskStop = iota
