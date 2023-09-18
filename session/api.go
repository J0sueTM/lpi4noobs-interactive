package session

import (
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/labstack/echo/v4"
)

func (*Session) Create(c echo.Context) error {
	return nil
}

const md string = `
# Opifex pectora vires mersa

## A Cnosiacas effugit

Lorem markdownum occuluere mihi aevi rerum in dicere meis refugit **noctes**.
Est praebet, extremas Parosque carpitur sacri committere, relevat quod? Hic ego
ex dumque de prima conscia relicta Ammon. Ante tantus et manifesta inque;
barbara et servatoremque iuvenis inobrutus tabula.

- Saturnia frondes et Cenchreis
- Letiferis tellure contermina verbis
- Emittite Parin aquae vincula et solidis Pisenore

## Sorores intabescere insculpunt egerit Dulichiae mediamque ante

Sequenti fretumque fortius vitiumque, tum retia quoque, non, per currus, tibi.
Sic possit potitur, pennis aut, te succiduo armigerae? Poteras et ferre
raptamque monstra deponere [utque erat](http://conspicuus.com/) suo imber.
Virilibus qui [Hermaphroditus quae](http://www.tibiquamvis.io/servabitur.php)
urguent tradit vocari lectus, alteraque avenis prementem radix.

## Penthea et utque ab inarsit occupat aquis

Vidit doctus potest sua **solida** electae, scissaque limus Cinyreius monte!
Aestus et dextra bicorni nec **rarus et** salire ad et detrahis mota restant ne
[texique Dauno](http://clamore-tela.org/negarein).

## Votum pelagi

Circumdata nec ultima munera ubi virtus nitorem possidet videres: ait *in scire
furta*, teneo sunt ille lentae. Infantem fortis, incinxit, ante latices legit.
Regia [alis](http://sit.net/utqueinconstantia.aspx) patiar gelidis salutifer
tumulo! Priami iamque! Studio Luna prole, consurgere interdum ausim Lycum Iubam
altissima pectoraque leonis Icare, visa.

1. Spectabat illa praeterea victa
2. Usus subduxit
3. Iuppiter AI viaeque
4. Quoque nate tantum viro totidem aures pectore
5. Quibus hac
6. Clipeum nec turba vobis illam

**Requirit** insidias nisi, qui restabam potuit vitarit atque, ferro flos illi
piae. Famem miratur pectore querenda, morte amnis, annos caelo
[et](http://adhuc-sua.net/). Ne eadem sonat et spectans iustis **non forte
seque**: uteri fuit cornu pressant.
`

func (*Session) Content(c echo.Context) error {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags: htmlFlags,
	}
	renderer := html.NewRenderer(opts)
	rndrDoc := markdown.Render(doc, renderer)

	return c.HTML(http.StatusOK, string(rndrDoc))
}
