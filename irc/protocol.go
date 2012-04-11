package irc

import (
    "strings"
)

// A single line from the IRC server
type Line struct {
    Nick, Ident, Host, Src string
	Cmd, Raw string
	Args     []string
}

// Parse a Line
func ParseLine(s string) *Line {
    line := &Line{Raw: s}
    if s[0] == ':' {
        if idx := strings.Index(s, " "); idx != -1 {
            line.Src, s = s[1:idx], s[idx+1:]
        } else {
            return nil
        }
    }
    args := strings.SplitN(s, " :", 2)
    if len(args) > 1 {
        args = append(strings.Fields(args[0]), args[1])
    } else {
        args = strings.Fields(args[0])
    }
    line.Cmd = strings.ToUpper(args[0])
    if len(args) > 1 {
        line.Args = args[1:]
    }
    return line
}





const (
    // Commands
    CMD_AWAY      = "AWAY"
    CMD_CLEARCHAN = "CLEARCHAN"
    CMD_CLEAROPS  = "CLEAROPS"
    CMD_DIE       = "DIE"
    CMD_ERROR     = "ERROR"
    CMD_HELP      = "HELP"
    CMD_INVITE    = "INVITE"
    CMD_JOIN      = "JOIN"
    CMD_KICK      = "KICK"
    CMD_KILL      = "KILL"
    CMD_KLINE     = "KLINE"
    CMD_LIST      = "LIST"
    CMD_MODE      = "MODE"
    CMD_NAMES     = "NAMES"
    CMD_NICK      = "NICK"
    CMD_NOTICE    = "NOTICE"
    CMD_OJOIN     = "OJOIN"
    CMD_OPER      = "OPER"
    CMD_OPME      = "OPME"
    CMD_PART      = "PART"
    CMD_PASS      = "PASS"
    CMD_PING      = "PING"
    CMD_PONG      = "PONG"
    CMD_PRIVMSG   = "PRIVMSG"
    CMD_QUIT      = "QUIT"
    CMD_STATS     = "STATS"
    CMD_SQUIT     = "SQUIT"
    CMD_TOPIC     = "TOPIC"
    CMD_USER      = "USER"
    CMD_WALLOPS   = "WALLOPS"
    CMD_WHO       = "WHO"
    CMD_WHOIS     = "WHOIS"

    // Signon responses
    RPL_WELCOME  = "001"
    RPL_YOURHOST = "002"
    RPL_CREATED  = "003"
    RPL_MYINFO   = "004"
    RPL_ISUPPORT = "005"

    // Command replies
    RPL_ADMINEMAIL      = "259"
    RPL_ADMINLOC1       = "257"
    RPL_ADMINLOC2       = "258"
    RPL_ADMINME         = "256"
    RPL_AWAY            = "301"
    RPL_BANLIST         = "367"
    RPL_CHANNELMODEIS   = "324"
    RPL_ENDOFBANLIST    = "368"
    RPL_ENDOFEXCEPTLIST = "349"
    RPL_ENDOFINFO       = "374"
    RPL_ENDOFINVITELIST = "347"
    RPL_ENDOFLINKS      = "365"
    RPL_ENDOFMOTD       = "376"
    RPL_ENDOFNAMES      = "366"
    RPL_ENDOFSTATS      = "219"
    RPL_ENDOFUSERS      = "394"
    RPL_ENDOFWHO        = "315"
    RPL_ENDOFWHOIS      = "318"
    RPL_EXCEPTLIST      = "348"
    RPL_INFO            = "371"
    RPL_INVITELIST      = "346"
    RPL_INVITING        = "341"
    RPL_ISON            = "303"
    RPL_LINKS           = "364"
    RPL_LISTSTART       = "321"
    RPL_LIST            = "322"
    RPL_LISTEND         = "323"
    RPL_LUSERCHANNELS   = "254"
    RPL_LUSERCLIENT     = "251"
    RPL_LUSERME         = "255"
    RPL_LUSEROP         = "252"
    RPL_LUSERUNKNOWN    = "253"
    RPL_MOTD            = "372"
    RPL_MOTDSTART       = "375"
    RPL_NAMEREPLY       = "353"
    RPL_NOTOPIC         = "331"
    RPL_NOUSERS         = "395"
    RPL_NOWAWAY         = "306"
    RPL_REHASHING       = "382"
    RPL_SERVLIST        = "234"
    RPL_SERVLISTEND     = "235"
    RPL_STATSCOMMANDS   = "212"
    RPL_STATSLINKINFO   = "211"
    RPL_STATSOLINE      = "243"
    RPL_STATSpLINE      = "249"
    RPL_STATSPLINE      = "220"
    RPL_STATSUPTIME     = "242"
    RPL_SUMMONING       = "342"
    RPL_TIME            = "391"
    RPL_TOPIC           = "332"
    RPL_TOPICWHOTIME    = "333"
    RPL_TRACECLASS      = "209"
    RPL_TRACECONNECTING = "201"
    RPL_TRACEEND        = "262"
    RPL_TRACEHANDSHAKE  = "202"
    RPL_TRACELINK       = "200"
    RPL_TRACELOG        = "261"
    RPL_TRACENEWTYPE    = "208"
    RPL_TRACEOPERATOR   = "204"
    RPL_TRACERECONNECT  = "210"
    RPL_TRACESERVER     = "206"
    RPL_TRACESERVICE    = "207"
    RPL_TRACEUNKNOWN    = "203"
    RPL_TRACEUSER       = "205"
    RPL_TRYAGAIN        = "263"
    RPL_UMODEIS         = "221"
    RPL_UNAWAY          = "305"
    RPL_UNIQOPIS        = "325"
    RPL_USERHOST        = "302"
    RPL_USERS           = "393"
    RPL_USERSSTART      = "392"
    RPL_VERSION         = "351"
    RPL_WHOISACCOUNT    = "330"
    RPL_WHOISCHANNELS   = "319"
    RPL_WHOISHOST       = "378"
    RPL_WHOISIDLE       = "317"
    RPL_WHOISMODES      = "379"
    RPL_WHOISOPERATOR   = "313"
    RPL_WHOISSECURE     = "671"
    RPL_WHOISSERVER     = "312"
    RPL_WHOISUSER       = "311"
    RPL_WHOREPLY        = "352"
    RPL_YOUREOPER       = "381"
    RPL_YOURESERVICE    = "383"

    // Error replies
    ERR_ALREADYREGISTERED = "462"
    ERR_BADCHANMASK       = "476"
    ERR_BADCHANNELKEY     = "475"
    ERR_BADMASK           = "415"
    ERR_BANLISTFULL       = "478"
    ERR_BANNEDFROMCHAN    = "474"
    ERR_CANNOTSENDTOCHAN  = "404"
    ERR_CANTKILLSERVER    = "483"
    ERR_CHANNELISFULL     = "471"
    ERR_CHANOPRIVSNEEDED  = "482"
    ERR_ERRONEOUSNICKNAME = "432"
    ERR_FILEERROR         = "424"
    ERR_INVITEONLYCHAN    = "473"
    ERR_KEYSET            = "467"
    ERR_NEEDMOREPARAMS    = "461"
    ERR_NICKCOLLISION     = "436"
    ERR_NICKNAMEINUSE     = "433"
    ERR_NOADMININFO       = "423"
    ERR_NOCHANMODES       = "477"
    ERR_NOLOGIN           = "444"
    ERR_NOMOTD            = "422"
    ERR_NONICKNAMEGIVEN   = "431"
    ERR_NOOPERHOST        = "491"
    ERR_NOORIGIN          = "409"
    ERR_NOPERMFORHOST     = "463"
    ERR_NOPRIVILEGES      = "481"
    ERR_NORECIPIENT       = "411"
    ERR_NOSUCHCHANNEL     = "403"
    ERR_NOSUCHNICK        = "401"
    ERR_NOSUCHSERVER      = "402"
    ERR_NOSUCHSERVICE     = "408"
    ERR_NOTEXTTOSEND      = "412"
    ERR_NOTONCHANNEL      = "442"
    ERR_NOTOPLEVEL        = "413"
    ERR_NOTREGISTERED     = "451"
    ERR_PASSWDMISMATCH    = "464"
    ERR_RESTRICTED        = "484"
    ERR_SUMMONDISABLED    = "445"
    ERR_TOOMANYCHANNELS   = "405"
    ERR_TOOMANYTARGETS    = "407"
    ERR_UMODEUNKNOWNFLAG  = "501"
    ERR_UNAVAILRESOURCE   = "437"
    ERR_UNIQOPPRIVSNEEDED = "485"
    ERR_UNKNOWNCOMMAND    = "421"
    ERR_UNKNOWNMODE       = "472"
    ERR_USERNOTINCHANNEL  = "441"
    ERR_USERONCHANNEL     = "443"
    ERR_USERSDISABLED     = "446"
    ERR_USERSDONTMATCH    = "502"
    ERR_WASNOSUCHNICK     = "406"
    ERR_WILDTOPLEVEL      = "414"
    ERR_YOUREBANNEDCREEP  = "465"
    ERR_YOUWILLBEBANNED   = "466"
)

// More

const (
    ERR_ILLEGALCHAN = "479"
)
