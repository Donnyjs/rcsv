package constant

const (
	INSCRIPTION_INFO           = "https://ordinals.com/content/%s"
	INSCRIPTION_LIST           = "https://api.hiro.so/ordinals/v1/inscriptions?limit=%d&offset=%d%s"
	INSCRIPTION_LIST_ARGS      = "&order=desc&order_by=genesis_block_height&mime_type=image%2Fsvg%2Bxml&from_number=13067644" //12205851
	INSCRIPTION_LIST_INIT_ARGS = "&to_number=13067644"
	DATA_CLCT                  = "doodinals"
	DATA_RCSV_IO               = "rcsv.io"
	FETCH_ALL_ARGS             = "&order=asc&order_by=genesis_block_height&mime_type=image/svg%sxml&mime_type=text/html&from_number=%d&to_number=%d"
	MIME_HTML                  = "text/html"
	MIME_SVG                   = "image/svg+xml"
	RECURSIVE_MONITOR_ARGS     = "&order=desc&order_by=genesis_block_height&mime_type=image/svg%2Bxml&mime_type=text/html&from_number=14796126"
)
