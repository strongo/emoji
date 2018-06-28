package emojis

var CategoryFlags = []string {
	"🏳️", // white flag, tags=[], aliases=[white_flag]
	"🏴", // black flag, tags=[], aliases=[black_flag]
	"🏁", // chequered flag, tags=[milestone finish], aliases=[checkered_flag]
	"🚩", // triangular flag, tags=[], aliases=[triangular_flag_on_post]
	"🏳️‍🌈", // rainbow flag, tags=[pride], aliases=[rainbow_flag]
	"🇦🇫", // Afghanistan, tags=[], aliases=[afghanistan]
	"🇦🇽", // Åland Islands, tags=[], aliases=[aland_islands]
	"🇦🇱", // Albania, tags=[], aliases=[albania]
	"🇩🇿", // Algeria, tags=[], aliases=[algeria]
	"🇦🇸", // American Samoa, tags=[], aliases=[american_samoa]
	"🇦🇩", // Andorra, tags=[], aliases=[andorra]
	"🇦🇴", // Angola, tags=[], aliases=[angola]
	"🇦🇮", // Anguilla, tags=[], aliases=[anguilla]
	"🇦🇶", // Antarctica, tags=[], aliases=[antarctica]
	"🇦🇬", // Antigua & Barbuda, tags=[], aliases=[antigua_barbuda]
	"🇦🇷", // Argentina, tags=[], aliases=[argentina]
	"🇦🇲", // Armenia, tags=[], aliases=[armenia]
	"🇦🇼", // Aruba, tags=[], aliases=[aruba]
	"🇦🇺", // Australia, tags=[], aliases=[australia]
	"🇦🇹", // Austria, tags=[], aliases=[austria]
	"🇦🇿", // Azerbaijan, tags=[], aliases=[azerbaijan]
	"🇧🇸", // Bahamas, tags=[], aliases=[bahamas]
	"🇧🇭", // Bahrain, tags=[], aliases=[bahrain]
	"🇧🇩", // Bangladesh, tags=[], aliases=[bangladesh]
	"🇧🇧", // Barbados, tags=[], aliases=[barbados]
	"🇧🇾", // Belarus, tags=[], aliases=[belarus]
	"🇧🇪", // Belgium, tags=[], aliases=[belgium]
	"🇧🇿", // Belize, tags=[], aliases=[belize]
	"🇧🇯", // Benin, tags=[], aliases=[benin]
	"🇧🇲", // Bermuda, tags=[], aliases=[bermuda]
	"🇧🇹", // Bhutan, tags=[], aliases=[bhutan]
	"🇧🇴", // Bolivia, tags=[], aliases=[bolivia]
	"🇧🇶", // Caribbean Netherlands, tags=[], aliases=[caribbean_netherlands]
	"🇧🇦", // Bosnia & Herzegovina, tags=[], aliases=[bosnia_herzegovina]
	"🇧🇼", // Botswana, tags=[], aliases=[botswana]
	"🇧🇷", // Brazil, tags=[], aliases=[brazil]
	"🇮🇴", // British Indian Ocean Territory, tags=[], aliases=[british_indian_ocean_territory]
	"🇻🇬", // British Virgin Islands, tags=[], aliases=[british_virgin_islands]
	"🇧🇳", // Brunei, tags=[], aliases=[brunei]
	"🇧🇬", // Bulgaria, tags=[], aliases=[bulgaria]
	"🇧🇫", // Burkina Faso, tags=[], aliases=[burkina_faso]
	"🇧🇮", // Burundi, tags=[], aliases=[burundi]
	"🇨🇻", // Cape Verde, tags=[], aliases=[cape_verde]
	"🇰🇭", // Cambodia, tags=[], aliases=[cambodia]
	"🇨🇲", // Cameroon, tags=[], aliases=[cameroon]
	"🇨🇦", // Canada, tags=[], aliases=[canada]
	"🇮🇨", // Canary Islands, tags=[], aliases=[canary_islands]
	"🇰🇾", // Cayman Islands, tags=[], aliases=[cayman_islands]
	"🇨🇫", // Central African Republic, tags=[], aliases=[central_african_republic]
	"🇹🇩", // Chad, tags=[], aliases=[chad]
	"🇨🇱", // Chile, tags=[], aliases=[chile]
	"🇨🇳", // China, tags=[china], aliases=[cn]
	"🇨🇽", // Christmas Island, tags=[], aliases=[christmas_island]
	"🇨🇨", // Cocos (Keeling) Islands, tags=[keeling], aliases=[cocos_islands]
	"🇨🇴", // Colombia, tags=[], aliases=[colombia]
	"🇰🇲", // Comoros, tags=[], aliases=[comoros]
	"🇨🇬", // Congo - Brazzaville, tags=[], aliases=[congo_brazzaville]
	"🇨🇩", // Congo - Kinshasa, tags=[], aliases=[congo_kinshasa]
	"🇨🇰", // Cook Islands, tags=[], aliases=[cook_islands]
	"🇨🇷", // Costa Rica, tags=[], aliases=[costa_rica]
	"🇨🇮", // Côte d’Ivoire, tags=[ivory], aliases=[cote_divoire]
	"🇭🇷", // Croatia, tags=[], aliases=[croatia]
	"🇨🇺", // Cuba, tags=[], aliases=[cuba]
	"🇨🇼", // Curaçao, tags=[], aliases=[curacao]
	"🇨🇾", // Cyprus, tags=[], aliases=[cyprus]
	"🇨🇿", // Czech Republic, tags=[], aliases=[czech_republic]
	"🇩🇰", // Denmark, tags=[], aliases=[denmark]
	"🇩🇯", // Djibouti, tags=[], aliases=[djibouti]
	"🇩🇲", // Dominica, tags=[], aliases=[dominica]
	"🇩🇴", // Dominican Republic, tags=[], aliases=[dominican_republic]
	"🇪🇨", // Ecuador, tags=[], aliases=[ecuador]
	"🇪🇬", // Egypt, tags=[], aliases=[egypt]
	"🇸🇻", // El Salvador, tags=[], aliases=[el_salvador]
	"🇬🇶", // Equatorial Guinea, tags=[], aliases=[equatorial_guinea]
	"🇪🇷", // Eritrea, tags=[], aliases=[eritrea]
	"🇪🇪", // Estonia, tags=[], aliases=[estonia]
	"🇪🇹", // Ethiopia, tags=[], aliases=[ethiopia]
	"🇪🇺", // European Union, tags=[], aliases=[eu european_union]
	"🇫🇰", // Falkland Islands, tags=[], aliases=[falkland_islands]
	"🇫🇴", // Faroe Islands, tags=[], aliases=[faroe_islands]
	"🇫🇯", // Fiji, tags=[], aliases=[fiji]
	"🇫🇮", // Finland, tags=[], aliases=[finland]
	"🇫🇷", // France, tags=[france french], aliases=[fr]
	"🇬🇫", // French Guiana, tags=[], aliases=[french_guiana]
	"🇵🇫", // French Polynesia, tags=[], aliases=[french_polynesia]
	"🇹🇫", // French Southern Territories, tags=[], aliases=[french_southern_territories]
	"🇬🇦", // Gabon, tags=[], aliases=[gabon]
	"🇬🇲", // Gambia, tags=[], aliases=[gambia]
	"🇬🇪", // Georgia, tags=[], aliases=[georgia]
	"🇩🇪", // Germany, tags=[flag germany], aliases=[de]
	"🇬🇭", // Ghana, tags=[], aliases=[ghana]
	"🇬🇮", // Gibraltar, tags=[], aliases=[gibraltar]
	"🇬🇷", // Greece, tags=[], aliases=[greece]
	"🇬🇱", // Greenland, tags=[], aliases=[greenland]
	"🇬🇩", // Grenada, tags=[], aliases=[grenada]
	"🇬🇵", // Guadeloupe, tags=[], aliases=[guadeloupe]
	"🇬🇺", // Guam, tags=[], aliases=[guam]
	"🇬🇹", // Guatemala, tags=[], aliases=[guatemala]
	"🇬🇬", // Guernsey, tags=[], aliases=[guernsey]
	"🇬🇳", // Guinea, tags=[], aliases=[guinea]
	"🇬🇼", // Guinea-Bissau, tags=[], aliases=[guinea_bissau]
	"🇬🇾", // Guyana, tags=[], aliases=[guyana]
	"🇭🇹", // Haiti, tags=[], aliases=[haiti]
	"🇭🇳", // Honduras, tags=[], aliases=[honduras]
	"🇭🇰", // Hong Kong SAR China, tags=[], aliases=[hong_kong]
	"🇭🇺", // Hungary, tags=[], aliases=[hungary]
	"🇮🇸", // Iceland, tags=[], aliases=[iceland]
	"🇮🇳", // India, tags=[], aliases=[india]
	"🇮🇩", // Indonesia, tags=[], aliases=[indonesia]
	"🇮🇷", // Iran, tags=[], aliases=[iran]
	"🇮🇶", // Iraq, tags=[], aliases=[iraq]
	"🇮🇪", // Ireland, tags=[], aliases=[ireland]
	"🇮🇲", // Isle of Man, tags=[], aliases=[isle_of_man]
	"🇮🇱", // Israel, tags=[], aliases=[israel]
	"🇮🇹", // Italy, tags=[italy], aliases=[it]
	"🇯🇲", // Jamaica, tags=[], aliases=[jamaica]
	"🇯🇵", // Japan, tags=[japan], aliases=[jp]
	"🎌", // crossed flags, tags=[], aliases=[crossed_flags]
	"🇯🇪", // Jersey, tags=[], aliases=[jersey]
	"🇯🇴", // Jordan, tags=[], aliases=[jordan]
	"🇰🇿", // Kazakhstan, tags=[], aliases=[kazakhstan]
	"🇰🇪", // Kenya, tags=[], aliases=[kenya]
	"🇰🇮", // Kiribati, tags=[], aliases=[kiribati]
	"🇽🇰", // Kosovo, tags=[], aliases=[kosovo]
	"🇰🇼", // Kuwait, tags=[], aliases=[kuwait]
	"🇰🇬", // Kyrgyzstan, tags=[], aliases=[kyrgyzstan]
	"🇱🇦", // Laos, tags=[], aliases=[laos]
	"🇱🇻", // Latvia, tags=[], aliases=[latvia]
	"🇱🇧", // Lebanon, tags=[], aliases=[lebanon]
	"🇱🇸", // Lesotho, tags=[], aliases=[lesotho]
	"🇱🇷", // Liberia, tags=[], aliases=[liberia]
	"🇱🇾", // Libya, tags=[], aliases=[libya]
	"🇱🇮", // Liechtenstein, tags=[], aliases=[liechtenstein]
	"🇱🇹", // Lithuania, tags=[], aliases=[lithuania]
	"🇱🇺", // Luxembourg, tags=[], aliases=[luxembourg]
	"🇲🇴", // Macau SAR China, tags=[], aliases=[macau]
	"🇲🇰", // Macedonia, tags=[], aliases=[macedonia]
	"🇲🇬", // Madagascar, tags=[], aliases=[madagascar]
	"🇲🇼", // Malawi, tags=[], aliases=[malawi]
	"🇲🇾", // Malaysia, tags=[], aliases=[malaysia]
	"🇲🇻", // Maldives, tags=[], aliases=[maldives]
	"🇲🇱", // Mali, tags=[], aliases=[mali]
	"🇲🇹", // Malta, tags=[], aliases=[malta]
	"🇲🇭", // Marshall Islands, tags=[], aliases=[marshall_islands]
	"🇲🇶", // Martinique, tags=[], aliases=[martinique]
	"🇲🇷", // Mauritania, tags=[], aliases=[mauritania]
	"🇲🇺", // Mauritius, tags=[], aliases=[mauritius]
	"🇾🇹", // Mayotte, tags=[], aliases=[mayotte]
	"🇲🇽", // Mexico, tags=[], aliases=[mexico]
	"🇫🇲", // Micronesia, tags=[], aliases=[micronesia]
	"🇲🇩", // Moldova, tags=[], aliases=[moldova]
	"🇲🇨", // Monaco, tags=[], aliases=[monaco]
	"🇲🇳", // Mongolia, tags=[], aliases=[mongolia]
	"🇲🇪", // Montenegro, tags=[], aliases=[montenegro]
	"🇲🇸", // Montserrat, tags=[], aliases=[montserrat]
	"🇲🇦", // Morocco, tags=[], aliases=[morocco]
	"🇲🇿", // Mozambique, tags=[], aliases=[mozambique]
	"🇲🇲", // Myanmar (Burma), tags=[burma], aliases=[myanmar]
	"🇳🇦", // Namibia, tags=[], aliases=[namibia]
	"🇳🇷", // Nauru, tags=[], aliases=[nauru]
	"🇳🇵", // Nepal, tags=[], aliases=[nepal]
	"🇳🇱", // Netherlands, tags=[], aliases=[netherlands]
	"🇳🇨", // New Caledonia, tags=[], aliases=[new_caledonia]
	"🇳🇿", // New Zealand, tags=[], aliases=[new_zealand]
	"🇳🇮", // Nicaragua, tags=[], aliases=[nicaragua]
	"🇳🇪", // Niger, tags=[], aliases=[niger]
	"🇳🇬", // Nigeria, tags=[], aliases=[nigeria]
	"🇳🇺", // Niue, tags=[], aliases=[niue]
	"🇳🇫", // Norfolk Island, tags=[], aliases=[norfolk_island]
	"🇲🇵", // Northern Mariana Islands, tags=[], aliases=[northern_mariana_islands]
	"🇰🇵", // North Korea, tags=[], aliases=[north_korea]
	"🇳🇴", // Norway, tags=[], aliases=[norway]
	"🇴🇲", // Oman, tags=[], aliases=[oman]
	"🇵🇰", // Pakistan, tags=[], aliases=[pakistan]
	"🇵🇼", // Palau, tags=[], aliases=[palau]
	"🇵🇸", // Palestinian Territories, tags=[], aliases=[palestinian_territories]
	"🇵🇦", // Panama, tags=[], aliases=[panama]
	"🇵🇬", // Papua New Guinea, tags=[], aliases=[papua_new_guinea]
	"🇵🇾", // Paraguay, tags=[], aliases=[paraguay]
	"🇵🇪", // Peru, tags=[], aliases=[peru]
	"🇵🇭", // Philippines, tags=[], aliases=[philippines]
	"🇵🇳", // Pitcairn Islands, tags=[], aliases=[pitcairn_islands]
	"🇵🇱", // Poland, tags=[], aliases=[poland]
	"🇵🇹", // Portugal, tags=[], aliases=[portugal]
	"🇵🇷", // Puerto Rico, tags=[], aliases=[puerto_rico]
	"🇶🇦", // Qatar, tags=[], aliases=[qatar]
	"🇷🇪", // Réunion, tags=[], aliases=[reunion]
	"🇷🇴", // Romania, tags=[], aliases=[romania]
	"🇷🇺", // Russia, tags=[russia], aliases=[ru]
	"🇷🇼", // Rwanda, tags=[], aliases=[rwanda]
	"🇧🇱", // St. Barthélemy, tags=[], aliases=[st_barthelemy]
	"🇸🇭", // St. Helena, tags=[], aliases=[st_helena]
	"🇰🇳", // St. Kitts & Nevis, tags=[], aliases=[st_kitts_nevis]
	"🇱🇨", // St. Lucia, tags=[], aliases=[st_lucia]
	"🇵🇲", // St. Pierre & Miquelon, tags=[], aliases=[st_pierre_miquelon]
	"🇻🇨", // St. Vincent & Grenadines, tags=[], aliases=[st_vincent_grenadines]
	"🇼🇸", // Samoa, tags=[], aliases=[samoa]
	"🇸🇲", // San Marino, tags=[], aliases=[san_marino]
	"🇸🇹", // São Tomé & Príncipe, tags=[], aliases=[sao_tome_principe]
	"🇸🇦", // Saudi Arabia, tags=[], aliases=[saudi_arabia]
	"🇸🇳", // Senegal, tags=[], aliases=[senegal]
	"🇷🇸", // Serbia, tags=[], aliases=[serbia]
	"🇸🇨", // Seychelles, tags=[], aliases=[seychelles]
	"🇸🇱", // Sierra Leone, tags=[], aliases=[sierra_leone]
	"🇸🇬", // Singapore, tags=[], aliases=[singapore]
	"🇸🇽", // Sint Maarten, tags=[], aliases=[sint_maarten]
	"🇸🇰", // Slovakia, tags=[], aliases=[slovakia]
	"🇸🇮", // Slovenia, tags=[], aliases=[slovenia]
	"🇸🇧", // Solomon Islands, tags=[], aliases=[solomon_islands]
	"🇸🇴", // Somalia, tags=[], aliases=[somalia]
	"🇿🇦", // South Africa, tags=[], aliases=[south_africa]
	"🇬🇸", // South Georgia & South Sandwich Islands, tags=[], aliases=[south_georgia_south_sandwich_islands]
	"🇰🇷", // South Korea, tags=[korea], aliases=[kr]
	"🇸🇸", // South Sudan, tags=[], aliases=[south_sudan]
	"🇪🇸", // Spain, tags=[spain], aliases=[es]
	"🇱🇰", // Sri Lanka, tags=[], aliases=[sri_lanka]
	"🇸🇩", // Sudan, tags=[], aliases=[sudan]
	"🇸🇷", // Suriname, tags=[], aliases=[suriname]
	"🇸🇿", // Swaziland, tags=[], aliases=[swaziland]
	"🇸🇪", // Sweden, tags=[], aliases=[sweden]
	"🇨🇭", // Switzerland, tags=[], aliases=[switzerland]
	"🇸🇾", // Syria, tags=[], aliases=[syria]
	"🇹🇼", // Taiwan, tags=[], aliases=[taiwan]
	"🇹🇯", // Tajikistan, tags=[], aliases=[tajikistan]
	"🇹🇿", // Tanzania, tags=[], aliases=[tanzania]
	"🇹🇭", // Thailand, tags=[], aliases=[thailand]
	"🇹🇱", // Timor-Leste, tags=[], aliases=[timor_leste]
	"🇹🇬", // Togo, tags=[], aliases=[togo]
	"🇹🇰", // Tokelau, tags=[], aliases=[tokelau]
	"🇹🇴", // Tonga, tags=[], aliases=[tonga]
	"🇹🇹", // Trinidad & Tobago, tags=[], aliases=[trinidad_tobago]
	"🇹🇳", // Tunisia, tags=[], aliases=[tunisia]
	"🇹🇷", // Turkey, tags=[turkey], aliases=[tr]
	"🇹🇲", // Turkmenistan, tags=[], aliases=[turkmenistan]
	"🇹🇨", // Turks & Caicos Islands, tags=[], aliases=[turks_caicos_islands]
	"🇹🇻", // Tuvalu, tags=[], aliases=[tuvalu]
	"🇺🇬", // Uganda, tags=[], aliases=[uganda]
	"🇺🇦", // Ukraine, tags=[], aliases=[ukraine]
	"🇦🇪", // United Arab Emirates, tags=[], aliases=[united_arab_emirates]
	"🇬🇧", // United Kingdom, tags=[flag british], aliases=[gb uk]
	"🇺🇸", // United States, tags=[flag united america], aliases=[us]
	"🇻🇮", // U.S. Virgin Islands, tags=[], aliases=[us_virgin_islands]
	"🇺🇾", // Uruguay, tags=[], aliases=[uruguay]
	"🇺🇿", // Uzbekistan, tags=[], aliases=[uzbekistan]
	"🇻🇺", // Vanuatu, tags=[], aliases=[vanuatu]
	"🇻🇦", // Vatican City, tags=[], aliases=[vatican_city]
	"🇻🇪", // Venezuela, tags=[], aliases=[venezuela]
	"🇻🇳", // Vietnam, tags=[], aliases=[vietnam]
	"🇼🇫", // Wallis & Futuna, tags=[], aliases=[wallis_futuna]
	"🇪🇭", // Western Sahara, tags=[], aliases=[western_sahara]
	"🇾🇪", // Yemen, tags=[], aliases=[yemen]
	"🇿🇲", // Zambia, tags=[], aliases=[zambia]
	"🇿🇼", // Zimbabwe, tags=[], aliases=[zimbabwe]
}
