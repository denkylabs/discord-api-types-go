package gateway

import (
	"github.com/denkylabs/discord-api-types-go/globals"
	"github.com/denkylabs/discord-api-types-go/payloads"
)

const GatewayVersion string = "10"

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-opcodes
type GatewayOpcodes uint8

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-opcodes
const (
	// A event was dispatched
	GatewayOpcodesDispatch GatewayOpcodes = iota
	// A heartbeat was received
	GatewayOpcodesHeartbeat
	// Starts a new session during the initial handshake
	GatewayOpcodesIdentify
	// Update the client's presence
	GatewayOpcodesPresenceUpdate
	// Used to join/leave or move between voice channels
	GatewayOpcodesVoiceStateUpdate
	// Resume a previous session that was disconnected
	GatewayOpcodesResume GatewayOpcodes = 6
	// You should attempt to reconnect and resume immediately
	GatewayOpcodesReconnect GatewayOpcodes = 7
	// Request information about offline guild members in a large guild
	GatewayOpcodesRequestGuildMembers GatewayOpcodes = 8
	// The session has been invalidated. You should reconnect and identify/resume accordingly
	GatewayOpcodesInvalidSession GatewayOpcodes = 9
	// Sent immediately after connecting, contains the heartbeat_interval to use
	GatewayOpcodesHello GatewayOpcodes = 10
	// 	Sent in response to receiving a heartbeat to acknowledge that it has been received
	GatewayOpcodesHeartbeatACK GatewayOpcodes = 11
)

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-close-event-codes
type GatewayCloseCodes uint16

// https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-close-event-codes
const (
	// We're not sure what went wrong. Try reconnecting?
	GatewayCloseCodesUnknownError GatewayCloseCodes = 4000
	// You sent an invalid Gateway opcode or an invalid payload for an opcode. Don't do that!
	// See https://discord.com/developers/docs/topics/gateway#payloads-and-opcodes
	GatewayCloseCodesUnknownOpcode GatewayCloseCodes = 4001
	// You sent an invalid payload to us. Don't do that!
	// See https://discord.com/developers/docs/topics/gateway#sending-payloads
	GatewayCloseCodesDecodeError GatewayCloseCodes = 4002
	// You sent us a payload prior to identifying
	// See https://discord.com/developers/docs/topics/gateway#identify
	GatewayCloseCodesNotAuthenticated GatewayCloseCodes = 4003
	// The account token sent with your identify payload is incorrect
	// See https://discord.com/developers/docs/topics/gateway#identify
	GatewayCloseCodesAuthenticationFailed GatewayCloseCodes = 4004
	// You sent more than one identify payload. Don't do that!
	GatewayCloseCodesAlreadyAuthenticated GatewayCloseCodes = 4005
	// The sequence sent when resuming the session was invalid. Reconnect and start a new session
	// See https://discord.com/developers/docs/topics/gateway#resume
	GatewayCloseCodesInvalidSeq GatewayCloseCodes = 4007
	// Woah nelly! You're sending payloads to us too quickly. Slow it down! You will be disconnected on receiving this
	GatewayCloseCodesRateLimited GatewayCloseCodes = 4008
	// Your session timed out. Reconnect and start a new one
	GatewayCloseCodesSessionTimedOut GatewayCloseCodes = 4009
	// You sent us an invalid shard when identifying
	// See https://discord.com/developers/docs/topics/gateway#sharding
	GatewayCloseCodesInvalidShard GatewayCloseCodes = 4010
	// The session would have handled too many guilds - you are required to shard your connection in order to connect
	// See https://discord.com/developers/docs/topics/gateway#sharding
	GatewayCloseCodesShardingRequired GatewayCloseCodes = 4011
	// You sent an invalid version for the gateway
	GatewayCloseCodesInvalidAPIVersion GatewayCloseCodes = 4012
	// You sent an invalid intent for a Gateway Intent. You may have incorrectly calculated the bitwise value
	GatewayCloseCodesInvalidIntents GatewayCloseCodes = 4013
	// You sent a disallowed intent for a Gateway Intent. You may have tried to specify an intent that you have not
	// enabled or are not whitelisted for
	GatewayCloseCodesDisallowedIntents GatewayCloseCodes = 4014
)

// https://discord.com/developers/docs/topics/gateway#list-of-intents
type Intent uint32

// https://discord.com/developers/docs/topics/gateway#list-of-intents
const (
	IntentGuilds                 Intent = 1 << 0
	IntentGuildMembers           Intent = 1 << 1
	IntentGuildBans              Intent = 1 << 2
	IntentGuildEmojisAndStickers Intent = 1 << 3
	IntentGuildIntegrations      Intent = 1 << 4
	IntentGuildWebhooks          Intent = 1 << 5
	IntentGuildInvites           Intent = 1 << 6
	IntentGuildVoiceStates       Intent = 1 << 7
	IntentGuildPresences         Intent = 1 << 8
	IntentGuildMessages          Intent = 1 << 9
	IntentGuildMessageReactions  Intent = 1 << 10
	IntentGuildMessageTyping     Intent = 1 << 11
	IntentDirectMessages         Intent = 1 << 12
	IntentDirectMessageReactions Intent = 1 << 13
	IntentDirectMessageTyping    Intent = 1 << 14
	IntentMessageContent         Intent = 1 << 15
	IntentGuildScheduledEvents   Intent = 1 << 16
)

// https://discord.com/developers/docs/topics/gateway#commands-and-events-gateway-events
type GatewayDispatchEvents string

// https://discord.com/developers/docs/topics/gateway#commands-and-events-gateway-events
const (
	GatewayDispatchEventsApplicationCommandPermissionsUpdate GatewayDispatchEvents = "APPLICATION_COMMAND_PERMISSIONS_UPDATE"
	GatewayDispatchEventsChannelCreate                       GatewayDispatchEvents = "CHANNEL_CREATE"
	GatewayDispatchEventsChannelDelete                       GatewayDispatchEvents = "CHANNEL_DELETE"
	GatewayDispatchEventsChannelPinsUpdate                   GatewayDispatchEvents = "CHANNEL_PINS_UPDATE"
	GatewayDispatchEventsChannelUpdate                       GatewayDispatchEvents = "CHANNEL_UPDATE"
	GatewayDispatchEventsGuildBanAdd                         GatewayDispatchEvents = "GUILD_BAN_ADD"
	GatewayDispatchEventsGuildBanRemove                      GatewayDispatchEvents = "GUILD_BAN_REMOVE"
	GatewayDispatchEventsGuildCreate                         GatewayDispatchEvents = "GUILD_CREATE"
	GatewayDispatchEventsGuildDelete                         GatewayDispatchEvents = "GUILD_DELETE"
	GatewayDispatchEventsGuildEmojisUpdate                   GatewayDispatchEvents = "GUILD_EMOJIS_UPDATE"
	GatewayDispatchEventsGuildIntegrationsUpdate             GatewayDispatchEvents = "GUILD_INTEGRATIONS_UPDATE"
	GatewayDispatchEventsGuildMemberAdd                      GatewayDispatchEvents = "GUILD_MEMBER_ADD"
	GatewayDispatchEventsGuildMemberRemove                   GatewayDispatchEvents = "GUILD_MEMBER_REMOVE"
	GatewayDispatchEventsGuildMembersChunk                   GatewayDispatchEvents = "GUILD_MEMBERS_CHUNK"
	GatewayDispatchEventsGuildMemberUpdate                   GatewayDispatchEvents = "GUILD_MEMBER_UPDATE"
	GatewayDispatchEventsGuildRoleCreate                     GatewayDispatchEvents = "GUILD_ROLE_CREATE"
	GatewayDispatchEventsGuildRoleDelete                     GatewayDispatchEvents = "GUILD_ROLE_DELETE"
	GatewayDispatchEventsGuildRoleUpdate                     GatewayDispatchEvents = "GUILD_ROLE_UPDATE"
	GatewayDispatchEventsGuildStickersUpdate                 GatewayDispatchEvents = "GUILD_STICKERS_UPDATE"
	GatewayDispatchEventsGuildUpdate                         GatewayDispatchEvents = "GUILD_UPDATE"
	GatewayDispatchEventsIntegrationCreate                   GatewayDispatchEvents = "INTEGRATION_CREATE"
	GatewayDispatchEventsIntegrationDelete                   GatewayDispatchEvents = "INTEGRATION_DELETE"
	GatewayDispatchEventsIntegrationUpdate                   GatewayDispatchEvents = "INTEGRATION_UPDATE"
	GatewayDispatchEventsInteractionCreate                   GatewayDispatchEvents = "INTERACTION_CREATE"
	GatewayDispatchEventsInviteCreate                        GatewayDispatchEvents = "INVITE_CREATE"
	GatewayDispatchEventsInviteDelete                        GatewayDispatchEvents = "INVITE_DELETE"
	GatewayDispatchEventsMessageCreate                       GatewayDispatchEvents = "MESSAGE_CREATE"
	GatewayDispatchEventsMessageDelete                       GatewayDispatchEvents = "MESSAGE_DELETE"
	GatewayDispatchEventsMessageDeleteBulk                   GatewayDispatchEvents = "MESSAGE_DELETE_BULK"
	GatewayDispatchEventsMessageReactionAdd                  GatewayDispatchEvents = "MESSAGE_REACTION_ADD"
	GatewayDispatchEventsMessageReactionRemove               GatewayDispatchEvents = "MESSAGE_REACTION_REMOVE"
	GatewayDispatchEventsMessageReactionRemoveAll            GatewayDispatchEvents = "MESSAGE_REACTION_REMOVE_ALL"
	GatewayDispatchEventsMessageReactionRemoveEmoji          GatewayDispatchEvents = "MESSAGE_REACTION_REMOVE_EMOJI"
	GatewayDispatchEventsMessageUpdate                       GatewayDispatchEvents = "MESSAGE_UPDATE"
	GatewayDispatchEventsPresenceUpdate                      GatewayDispatchEvents = "PRESENCE_UPDATE"
	GatewayDispatchEventsStageInstanceCreate                 GatewayDispatchEvents = "STAGE_INSTANCE_CREATE"
	GatewayDispatchEventsStageInstanceDelete                 GatewayDispatchEvents = "STAGE_INSTANCE_DELETE"
	GatewayDispatchEventsStageInstanceUpdate                 GatewayDispatchEvents = "STAGE_INSTANCE_UPDATE"
	GatewayDispatchEventsReady                               GatewayDispatchEvents = "READY"
	GatewayDispatchEventsResumed                             GatewayDispatchEvents = "RESUMED"
	GatewayDispatchEventsThreadCreate                        GatewayDispatchEvents = "THREAD_CREATE"
	GatewayDispatchEventsThreadDelete                        GatewayDispatchEvents = "THREAD_DELETE"
	GatewayDispatchEventsThreadListSync                      GatewayDispatchEvents = "THREAD_LIST_SYNC"
	GatewayDispatchEventsThreadMembersUpdate                 GatewayDispatchEvents = "THREAD_MEMBERS_UPDATE"
	GatewayDispatchEventsThreadMemberUpdate                  GatewayDispatchEvents = "THREAD_MEMBER_UPDATE"
	GatewayDispatchEventsThreadUpdate                        GatewayDispatchEvents = "THREAD_UPDATE"
	GatewayDispatchEventsTypingStart                         GatewayDispatchEvents = "TYPING_START"
	GatewayDispatchEventsUserUpdate                          GatewayDispatchEvents = "USER_UPDATE"
	GatewayDispatchEventsVoiceServerUpdate                   GatewayDispatchEvents = "VOICE_SERVER_UPDATE"
	GatewayDispatchEventsVoiceStateUpdate                    GatewayDispatchEvents = "VOICE_STATE_UPDATE"
	GatewayDispatchEventsWebhooksUpdate                      GatewayDispatchEvents = "WEBHOOKS_UPDATE"
	GatewayDispatchEventsGuildScheduledEventCreate           GatewayDispatchEvents = "GUILD_SCHEDULED_EVENT_CREATE"
	GatewayDispatchEventsGuildScheduledEventUpdate           GatewayDispatchEvents = "GUILD_SCHEDULED_EVENT_UPDATE"
	GatewayDispatchEventsGuildScheduledEventDelete           GatewayDispatchEvents = "GUILD_SCHEDULED_EVENT_DELETE"
	GatewayDispatchEventsGuildScheduledEventUserAdd          GatewayDispatchEvents = "GUILD_SCHEDULED_EVENT_USER_ADD"
	GatewayDispatchEventsGuildScheduledEventUserRemove       GatewayDispatchEvents = "GUILD_SCHEDULED_EVENT_USER_REMOVE"
)

// Gateway Dispatch Payloads

// https://discord.com/developers/docs/topics/gateway#hello
type GatewayHello NonDispatchPayload[GatewayHelloData]

// https://discord.com/developers/docs/topics/gateway#hello
type GatewayHelloData struct {
	// The interval (in milliseconds) the client should heartbeat with
	HeartbeatInterval uint16 `json:"heartbeat_interval"`
}

// https://discord.com/developers/docs/topics/gateway#heartbeating
type GatewayHeartbeatRequest NonDispatchPayload[interface{}]

// https://discord.com/developers/docs/topics/gateway#heartbeating-example-gateway-heartbeat-ack
type GatewayHeartbeatAck NonDispatchPayload[interface{}]

// https://discord.com/developers/docs/topics/gateway#invalid-session
type GatewayInvalidSession NonDispatchPayload[GatewayInvalidSessionData]

// https://discord.com/developers/docs/topics/gateway#invalid-session
type GatewayInvalidSessionData bool

// https://discord.com/developers/docs/topics/gateway#reconnect
type GatewayReconnect NonDispatchPayload[interface{}]

// https://discord.com/developers/docs/topics/gateway#ready
type GatewayReadyDispatch DataPayload[GatewayReadyDispatchData]

// https://discord.com/developers/docs/topics/gateway#ready
type GatewayReadyDispatchData struct {
	// Gateway version.
	// See https://discord.com/developers/docs/topics/gateway#gateways-gateway-versions
	V uint32 `json:"v"`
	// Information about the user including email
	// See https://discord.com/developers/docs/resources/user#user-object
	User payloads.APIUser `json:"user"`
	// The guilds the user is in
	// See https://discord.com/developers/docs/resources/guild#unavailable-guild-object
	Guilds payloads.APIUnavailableGuild `json:"guilds"`
	// Used for resuming connections
	SessionId string `json:"session_id"`
	// The shard information associated with this session, if sent when identifying. See https://discord.com/developers/docs/topics/gateway#sharding
	// See https://discord.com/developers/docs/topics/gateway#sharding
	Shard [2]uint32 `json:"shard"`
	// Contains 'id' and 'flags' properties
	// See https://discord.com/developers/docs/resources/application#application-object
	Application payloads.APIApplication `json:"application"`
}

// https://discord.com/developers/docs/topics/gateway#resumed
type GatewayResumedDispatch DataPayload[interface{}]

type GatewayChannelModifyDispatch DataPayload[GatewayChannelModifyDispatchData]

type GatewayChannelModifyDispatchData interface{} // TODO: Implement APIChannel type

// https://discord.com/developers/docs/topics/gateway#channel-update
type GatewayChannelCreateDispatch GatewayChannelModifyDispatch

// https://discord.com/developers/docs/topics/gateway#channel-update
type GatewayChannelCreateDispatchData GatewayChannelModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#channel-update
type GatewayChannelUpdateDispatch GatewayChannelModifyDispatch

// https://discord.com/developers/docs/topics/gateway#channel-update
type GatewayChannelUpdateDispatchData GatewayChannelModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#channel-update
type GatewayChannelDeleteDispatch GatewayChannelModifyDispatch

// https://discord.com/developers/docs/topics/gateway#channel-update
type GatewayChannelDeleteDispatchData GatewayChannelModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#channel-pins-update
type GatewayChannelPinsUpdateDispatch DataPayload[GatewayChannelPinsUpdateDispatchData]

// https://discord.com/developers/docs/topics/gateway#channel-pins-update
type GatewayChannelPinsUpdateDispatchData struct {
	// The id of the guild
	GuildId globals.Snowflake `json:"guild_id"`
	// The id of the channel
	ChannelId globals.Snowflake `json:"channel_id"`
	// The time at which the most recent pinned message was pinned
	LastPinTimestamp uint `json:"last_pin_timestamp"`
}

// https://discord.com/developers/docs/topics/gateway#guild-create
type GatewayGuildModifyDispatch DataPayload[GatewayGuildModifyDispatchData]

// https://discord.com/developers/docs/topics/gateway#guild-update
type GatewayGuildModifyDispatchData payloads.APIGuild

// https://discord.com/developers/docs/topics/gateway#guild-create
type GatewayGuildCreateDispatch DataPayload[GatewayGuildCreateDispatchData]

// https://discord.com/developers/docs/topics/gateway#guild-create
type GatewayGuildCreateDispatchData payloads.APIGuild

// https://discord.com/developers/docs/topics/gateway#guild-update
type GatewayGuildUpdateDispatch GatewayGuildModifyDispatch

// https://discord.com/developers/docs/topics/gateway#guild-update
type GatewayGuildUpdateDispatchData GatewayGuildModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#guild-delete
type GatewayGuildDeleteDispatch DataPayload[GatewayGuildDeleteDispatchData]

// https://discord.com/developers/docs/topics/gateway#guild-delete
type GatewayGuildDeleteDispatchData payloads.APIUnavailableGuild

// https://discord.com/developers/docs/topics/gateway#guild-ban-add
// https://discord.com/developers/docs/topics/gateway#guild-ban-remove
type GatewayGuildBanModifyDispatch DataPayload[GatewayGuildBanModifyDispatchData]

// https://discord.com/developers/docs/topics/gateway#guild-ban-add
// https://discord.com/developers/docs/topics/gateway#guild-ban-remove
type GatewayGuildBanModifyDispatchData struct {
	// ID of the guild
	GuildId globals.Snowflake `json:"guild_id"`
	// The banned user
	// See https://discord.com/developers/docs/resources/user#user-object
	User payloads.APIUser `json:"user"`
}

// https://discord.com/developers/docs/topics/gateway#guild-ban-add
type GatewayGuildBanAddDispatch GatewayGuildBanModifyDispatch

// https://discord.com/developers/docs/topics/gateway#guild-ban-add
type GatewayGuildBanAddDispatchData GatewayGuildBanModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#guild-ban-remove
type GatewayGuildBanRemoveDispatch GatewayGuildBanModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#guild-ban-remove
type GatewayGuildBanRemoveDispatchData GatewayGuildBanModifyDispatchData

// https://discord.com/developers/docs/topics/gateway#guild-emojis-update
type GatewayGuildEmojisUpdateDispatch DataPayload[GatewayGuildEmojisUpdateDispatchData]

// https://discord.com/developers/docs/topics/gateway#guild-emojis-update
type GatewayGuildEmojisUpdateDispatchData struct {
	// ID of the guild
	GuildId globals.Snowflake `json:"guild_id"`
	// Array of emojis
	// See https://discord.com/developers/docs/resources/emoji#emoji-object
	Emojis []payloads.APIEmoji `json:"emojis"`
}

// https://discord.com/developers/docs/topics/gateway#guild-stickers-update
type GatewayGuildStickersUpdateDispatch DataPayload[GatewayGuildStickersUpdateDispatchData]

// https://discord.com/developers/docs/topics/gateway#guild-stickers-update
type GatewayGuildStickersUpdateDispatchData struct {
	// ID of the guild
	GuildId globals.Snowflake `json:"guild_id"`
	// Array of stickers
	// See https://discord.com/developers/docs/resources/sticker#sticker-object
	Stickers []payloads.APISticker `json:"stickers"`
}

// https://discord.com/developers/docs/topics/gateway#guild-integrations-update
type GatewayGuildIntegrationsUpdateDispatch DataPayload[GatewayGuildIntegrationsUpdateDispatchData]

// https://discord.com/developers/docs/topics/gateway#guild-integrations-update
type GatewayGuildIntegrationsUpdateDispatchData struct {
	/// ID of the guild whose integrations were updated
	GuildId globals.Snowflake `json:"guild_id"`
}

// https://discord.com/developers/docs/topics/gateway#guild-member-add
type GatewayGuildMemberAddDispatch DataPayload[GatewayGuildMemberAddDispatchData]

// https://discord.com/developers/docs/topics/gateway#guild-member-add
type GatewayGuildMemberAddDispatchData struct {
	// The id of the guild
	GuildId globals.Snowflake `json:"guild_id"`
	// The user this guild member represents
	// **This field won't be included in the member object attached to `MESSAGE_CREATE` and `MESSAGE_UPDATE` gateway events.**
	// See https://discord.com/developers/docs/resources/user#user-object
	User payloads.APIUser `json:"user"`
	// This users guild nickname
	Nick string `json:"nick"`
	// The member's guild avatar hash
	Avatar string `json:"avatar"`
	// Array of role object ids
	// See https://discord.com/developers/docs/topics/permissions#role-object
	Roles []globals.Snowflake `json:"roles"`
	// When the user joined the guild
	JoinedAT GatewayDispatchEvents `json:"joined_at"`
	// When the user started boosting the guild
	// See https://support.discord.com/hc/en-us/articles/360028038352-Server-Boosting-
	PremiumSince string `json:"premium_since"`
	// Whether the user is deafened in voice channels
	Deaf bool `json:"deaf"`
	// Whether the user is muted in voice channels
	Mute bool `json:"mute"`
	// Whether the user has not yet passed the guild's Membership Screening requirements
	// If this field is not present, it can be assumed as `false`.*
	Pending bool `json:"pending"`
	// Timestamp of when the time out will be removed until then, they cannot interact with the guild
	CommunicationDisabledUntil string `json:"communication_disabled_until"`
}

// https://discord.com/developers/docs/topics/gateway#guild-member-remove
type GatewayGuildMemberRemoveDispatch DataPayload[GatewayGuildMemberRemoveDispatchData]

// https://discord.com/developers/docs/topics/gateway#guild-member-remove
type GatewayGuildMemberRemoveDispatchData struct {
	// The id of the guild
	GuildId globals.Snowflake `json:"guild_id"`
	// The user who was removed
	// See https://discord.com/developers/docs/resources/user#user-object
	User payloads.APIUser `json:"user"`
}

type NonDispatchPayload[D interface{}] struct {
	// Opcode for the payload
	Op uint8 `json:"op"`
	// Event data. Can sometimes be nil
	D D `json:"d"`
	// Always nil
	T interface{} `json:"t"`
	// Always nil
	S interface{} `json:"s"`
}

type DataPayload[D interface{}] struct {
	// Opcode for the payload (in this case will always be 0)
	Op GatewayOpcodes `json:"op"`
	// Event data
	D D `json:"d"`
	// Sequence number, used for resuming sessions and heartbeats
	S uint32 `json:"s"`
	// The event name for this payload
	T GatewayDispatchEvents `json:"t"`
}
