package emoji

var CategoryObjects = []string {
	"⌚️", // watch, tags=[time], aliases=[watch]
	"📱", // mobile phone, tags=[smartphone mobile], aliases=[iphone]
	"📲", // mobile phone with arrow, tags=[call incoming], aliases=[calling]
	"💻", // laptop computer, tags=[desktop screen], aliases=[computer]
	"⌨️", // keyboard, tags=[], aliases=[keyboard]
	"🖥", // desktop computer, tags=[], aliases=[desktop_computer]
	"🖨", // printer, tags=[], aliases=[printer]
	"🖱", // computer mouse, tags=[], aliases=[computer_mouse]
	"🖲", // trackball, tags=[], aliases=[trackball]
	"🕹", // joystick, tags=[], aliases=[joystick]
	"🗜", // clamp, tags=[], aliases=[clamp]
	"💽", // computer disk, tags=[], aliases=[minidisc]
	"💾", // floppy disk, tags=[save], aliases=[floppy_disk]
	"💿", // optical disk, tags=[], aliases=[cd]
	"📀", // dvd, tags=[], aliases=[dvd]
	"📼", // videocassette, tags=[], aliases=[vhs]
	"📷", // camera, tags=[photo], aliases=[camera]
	"📸", // camera with flash, tags=[photo], aliases=[camera_flash]
	"📹", // video camera, tags=[], aliases=[video_camera]
	"🎥", // movie camera, tags=[film video], aliases=[movie_camera]
	"📽", // film projector, tags=[], aliases=[film_projector]
	"🎞", // film frames, tags=[], aliases=[film_strip]
	"📞", // telephone receiver, tags=[phone call], aliases=[telephone_receiver]
	"☎️", // telephone, tags=[], aliases=[phone telephone]
	"📟", // pager, tags=[], aliases=[pager]
	"📠", // fax machine, tags=[], aliases=[fax]
	"📺", // television, tags=[], aliases=[tv]
	"📻", // radio, tags=[podcast], aliases=[radio]
	"🎙", // studio microphone, tags=[podcast], aliases=[studio_microphone]
	"🎚", // level slider, tags=[], aliases=[level_slider]
	"🎛", // control knobs, tags=[], aliases=[control_knobs]
	"⏱", // stopwatch, tags=[], aliases=[stopwatch]
	"⏲", // timer clock, tags=[], aliases=[timer_clock]
	"⏰", // alarm clock, tags=[morning], aliases=[alarm_clock]
	"🕰", // mantelpiece clock, tags=[], aliases=[mantelpiece_clock]
	"⌛️", // hourglass, tags=[time], aliases=[hourglass]
	"⏳", // hourglass with flowing sand, tags=[time], aliases=[hourglass_flowing_sand]
	"📡", // satellite antenna, tags=[signal], aliases=[satellite]
	"🔋", // battery, tags=[power], aliases=[battery]
	"🔌", // electric plug, tags=[], aliases=[electric_plug]
	"💡", // light bulb, tags=[idea light], aliases=[bulb]
	"🔦", // flashlight, tags=[], aliases=[flashlight]
	"🕯", // candle, tags=[], aliases=[candle]
	"🗑", // wastebasket, tags=[trash], aliases=[wastebasket]
	"🛢", // oil drum, tags=[], aliases=[oil_drum]
	"💸", // money with wings, tags=[dollar], aliases=[money_with_wings]
	"💵", // dollar banknote, tags=[money], aliases=[dollar]
	"💴", // yen banknote, tags=[], aliases=[yen]
	"💶", // euro banknote, tags=[], aliases=[euro]
	"💷", // pound banknote, tags=[], aliases=[pound]
	"💰", // money bag, tags=[dollar cream], aliases=[moneybag]
	"💳", // credit card, tags=[subscription], aliases=[credit_card]
	"💎", // gem stone, tags=[diamond], aliases=[gem]
	"⚖️", // balance scale, tags=[], aliases=[balance_scale]
	"🔧", // wrench, tags=[tool], aliases=[wrench]
	"🔨", // hammer, tags=[tool], aliases=[hammer]
	"⚒", // hammer and pick, tags=[], aliases=[hammer_and_pick]
	"🛠", // hammer and wrench, tags=[], aliases=[hammer_and_wrench]
	"⛏", // pick, tags=[], aliases=[pick]
	"🔩", // nut and bolt, tags=[], aliases=[nut_and_bolt]
	"⚙️", // gear, tags=[], aliases=[gear]
	"⛓", // chains, tags=[], aliases=[chains]
	"🔫", // pistol, tags=[shoot weapon], aliases=[gun]
	"💣", // bomb, tags=[boom], aliases=[bomb]
	"🔪", // kitchen knife, tags=[cut chop], aliases=[hocho knife]
	"🗡", // dagger, tags=[], aliases=[dagger]
	"⚔️", // crossed swords, tags=[], aliases=[crossed_swords]
	"🛡", // shield, tags=[], aliases=[shield]
	"🚬", // cigarette, tags=[cigarette], aliases=[smoking]
	"⚰️", // coffin, tags=[funeral], aliases=[coffin]
	"⚱️", // funeral urn, tags=[], aliases=[funeral_urn]
	"🏺", // amphora, tags=[], aliases=[amphora]
	"🔮", // crystal ball, tags=[fortune], aliases=[crystal_ball]
	"📿", // prayer beads, tags=[], aliases=[prayer_beads]
	"💈", // barber pole, tags=[], aliases=[barber]
	"⚗️", // alembic, tags=[], aliases=[alembic]
	"🔭", // telescope, tags=[], aliases=[telescope]
	"🔬", // microscope, tags=[science laboratory investigate], aliases=[microscope]
	"🕳", // hole, tags=[], aliases=[hole]
	"💊", // pill, tags=[health medicine], aliases=[pill]
	"💉", // syringe, tags=[health hospital needle], aliases=[syringe]
	"🌡", // thermometer, tags=[], aliases=[thermometer]
	"🚽", // toilet, tags=[wc], aliases=[toilet]
	"🚰", // potable water, tags=[], aliases=[potable_water]
	"🚿", // shower, tags=[bath], aliases=[shower]
	"🛁", // bathtub, tags=[], aliases=[bathtub]
	"🛀", // person taking bath, tags=[shower], aliases=[bath]
	"🛎", // bellhop bell, tags=[], aliases=[bellhop_bell]
	"🔑", // key, tags=[lock password], aliases=[key]
	"🗝", // old key, tags=[], aliases=[old_key]
	"🚪", // door, tags=[], aliases=[door]
	"🛋", // couch and lamp, tags=[], aliases=[couch_and_lamp]
	"🛏", // bed, tags=[], aliases=[bed]
	"🛌", // person in bed, tags=[], aliases=[sleeping_bed]
	"🖼", // framed picture, tags=[], aliases=[framed_picture]
	"🛍", // shopping bags, tags=[bags], aliases=[shopping]
	"🛒", // shopping cart, tags=[], aliases=[shopping_cart]
	"🎁", // wrapped gift, tags=[present birthday christmas], aliases=[gift]
	"🎈", // balloon, tags=[party birthday], aliases=[balloon]
	"🎏", // carp streamer, tags=[], aliases=[flags]
	"🎀", // ribbon, tags=[], aliases=[ribbon]
	"🎊", // confetti ball, tags=[], aliases=[confetti_ball]
	"🎉", // party popper, tags=[hooray party], aliases=[tada]
	"🎎", // Japanese dolls, tags=[], aliases=[dolls]
	"🏮", // red paper lantern, tags=[], aliases=[izakaya_lantern lantern]
	"🎐", // wind chime, tags=[], aliases=[wind_chime]
	"✉️", // envelope, tags=[letter], aliases=[email envelope]
	"📩", // envelope with arrow, tags=[], aliases=[envelope_with_arrow]
	"📨", // incoming envelope, tags=[], aliases=[incoming_envelope]
	"📧", // e-mail, tags=[], aliases=[e-mail]
	"💌", // love letter, tags=[email envelope], aliases=[love_letter]
	"📥", // inbox tray, tags=[], aliases=[inbox_tray]
	"📤", // outbox tray, tags=[], aliases=[outbox_tray]
	"📦", // package, tags=[shipping], aliases=[package]
	"🏷", // label, tags=[tag], aliases=[label]
	"📪", // closed mailbox with lowered flag, tags=[], aliases=[mailbox_closed]
	"📫", // closed mailbox with raised flag, tags=[], aliases=[mailbox]
	"📬", // open mailbox with raised flag, tags=[], aliases=[mailbox_with_mail]
	"📭", // open mailbox with lowered flag, tags=[], aliases=[mailbox_with_no_mail]
	"📮", // postbox, tags=[], aliases=[postbox]
	"📯", // postal horn, tags=[], aliases=[postal_horn]
	"📜", // scroll, tags=[document], aliases=[scroll]
	"📃", // page with curl, tags=[], aliases=[page_with_curl]
	"📄", // page facing up, tags=[document], aliases=[page_facing_up]
	"📑", // bookmark tabs, tags=[], aliases=[bookmark_tabs]
	"📊", // bar chart, tags=[stats metrics], aliases=[bar_chart]
	"📈", // chart increasing, tags=[graph metrics], aliases=[chart_with_upwards_trend]
	"📉", // chart decreasing, tags=[graph metrics], aliases=[chart_with_downwards_trend]
	"🗒", // spiral notepad, tags=[], aliases=[spiral_notepad]
	"🗓", // spiral calendar, tags=[], aliases=[spiral_calendar]
	"📆", // tear-off calendar, tags=[schedule], aliases=[calendar]
	"📅", // calendar, tags=[calendar schedule], aliases=[date]
	"📇", // card index, tags=[], aliases=[card_index]
	"🗃", // card file box, tags=[], aliases=[card_file_box]
	"🗳", // ballot box with ballot, tags=[], aliases=[ballot_box]
	"🗄", // file cabinet, tags=[], aliases=[file_cabinet]
	"📋", // clipboard, tags=[], aliases=[clipboard]
	"📁", // file folder, tags=[directory], aliases=[file_folder]
	"📂", // open file folder, tags=[], aliases=[open_file_folder]
	"🗂", // card index dividers, tags=[], aliases=[card_index_dividers]
	"🗞", // rolled-up newspaper, tags=[press], aliases=[newspaper_roll]
	"📰", // newspaper, tags=[press], aliases=[newspaper]
	"📓", // notebook, tags=[], aliases=[notebook]
	"📔", // notebook with decorative cover, tags=[], aliases=[notebook_with_decorative_cover]
	"📒", // ledger, tags=[], aliases=[ledger]
	"📕", // closed book, tags=[], aliases=[closed_book]
	"📗", // green book, tags=[], aliases=[green_book]
	"📘", // blue book, tags=[], aliases=[blue_book]
	"📙", // orange book, tags=[], aliases=[orange_book]
	"📚", // books, tags=[library], aliases=[books]
	"📖", // open book, tags=[], aliases=[book open_book]
	"🔖", // bookmark, tags=[], aliases=[bookmark]
	"🔗", // link, tags=[], aliases=[link]
	"📎", // paperclip, tags=[], aliases=[paperclip]
	"🖇", // linked paperclips, tags=[], aliases=[paperclips]
	"📐", // triangular ruler, tags=[], aliases=[triangular_ruler]
	"📏", // straight ruler, tags=[], aliases=[straight_ruler]
	"📌", // pushpin, tags=[location], aliases=[pushpin]
	"📍", // round pushpin, tags=[location], aliases=[round_pushpin]
	"✂️", // scissors, tags=[cut], aliases=[scissors]
	"🖊", // pen, tags=[], aliases=[pen]
	"🖋", // fountain pen, tags=[], aliases=[fountain_pen]
	"✒️", // black nib, tags=[], aliases=[black_nib]
	"🖌", // paintbrush, tags=[], aliases=[paintbrush]
	"🖍", // crayon, tags=[], aliases=[crayon]
	"📝", // memo, tags=[document note], aliases=[memo pencil]
	"✏️", // pencil, tags=[], aliases=[pencil2]
	"🔍", // left-pointing magnifying glass, tags=[search zoom], aliases=[mag]
	"🔎", // right-pointing magnifying glass, tags=[], aliases=[mag_right]
	"🔏", // locked with pen, tags=[], aliases=[lock_with_ink_pen]
	"🔐", // locked with key, tags=[security], aliases=[closed_lock_with_key]
	"🔒", // locked, tags=[security private], aliases=[lock]
	"🔓", // unlocked, tags=[security], aliases=[unlock]
}