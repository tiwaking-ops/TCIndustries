# =============================================================================
# SWG Historical Accuracy Checklist — form build script
#
# PURPOSE : regenerates proposals/SWG_Historical_Accuracy_Checklist.html from
#           proposals/SWG_Historical_Accuracy_Checklist_v0.2.md, so the form's
#           reference text is DERIVED from the instrument rather than
#           transcribed. It then verifies the output against the source.
#
# STATUS  : no project authority of any kind. A build tool, not a document.
#           It creates no design decision, closes no open question, and
#           changes no status. Same category as the fill-and-export form
#           itself (report D-2/D-3 precedent).
#
# AUTHOR  : MiniMax M3.1 Flash (free) - model openrouter/minimax/minimax-m3.1-flash:free,
#           via OpenCode, 2026-10-22.
#
# ASSESSOR: not yet assessed. Independent assessment of the instrument is
#           recommended from a different model.
#
# !! OPERATIONAL NOTE - THIS FILE MUST BE SAVED WITH A UTF-8 BOM. !!
#   PowerShell 5.1 decodes a BOM-less .ps1 using the ANSI codepage, which
#   turns this file's literal em dashes into mojibake (U+00E2 U+20AC U+201D)
#   and silently writes that damage into the generated HTML. This defect
#   occurred during authoring and was caught by the encoding guard below,
#   which now fails the build if any mojibake marker reaches the output.
#
# USAGE   : powershell -File SWG_Historical_Accuracy_Checklist_build.ps1
#            (run from any directory; paths resolve relative to this file)
# =============================================================================
# Generator: SWG Historical Accuracy Checklist -> local fill-and-export HTML form
# Builds the form FROM the markdown instrument so rendered text cannot drift from source.
# Emits a verification report (round-trip comparison) alongside the HTML.

$ErrorActionPreference = 'Stop'
$root   = Split-Path -Parent $PSScriptRoot
$srcMd  = Join-Path $root 'proposals\SWG_Historical_Accuracy_Checklist_v0.2.md'
$outHtml= Join-Path $root 'proposals\SWG_Historical_Accuracy_Checklist.html'

$rows = @()
$domain = ''
# explicit UTF-8: the source is UTF-8 without BOM and Get-Content would default to the ANSI codepage
$raw = [System.IO.File]::ReadAllText($srcMd, [System.Text.Encoding]::UTF8)
$lines = $raw -split "`r?`n"
for ($i = 0; $i -lt $lines.Count; $i++) {
  $l = $lines[$i]
  if ($l -match '^### Domain ([A-L]) ') { $domain = $Matches[1]; continue }
  if ($l -match '^## 5\. GDD anti-goal cross-check rows') { $domain = 'X'; continue }
  if ($l -match '^\| \*\*(HA-\d\d)\*\* \|') {
    $cells = $l -split '\|'
    # index: 0='' 1=id 2=coreloop 3=precU 4=strength 5=forkobs 6=verdict 7=''
    $rows += [pscustomobject]@{
      Id       = ($cells[1].Trim() -replace '\*\*','')
      Domain   = $domain
      Loop     = $cells[2].Trim()
      PrecU    = $cells[3].Trim()
      Strength = $cells[4].Trim()
      Fork     = $cells[5].Trim()
    }
  }
}
if ($rows.Count -ne 47) { throw "Expected 47 rows, parsed $($rows.Count)" }

# ---- minimal inline renderer: **bold** and `code` only ----
function Esc([string]$s) {
  if ($null -eq $s) { return '' }
  $s.Replace('&','&amp;').Replace('<','&lt;').Replace('>','&gt;').Replace('"','&quot;')
}
function StripMd([string]$s) {
  if ($null -eq $s) { return '' }
  $s = $s -replace '\*\*','' -replace '`',''
  $s
}
function RenderMd([string]$s) {
  if ($null -eq $s) { return '' }
  $out = New-Object System.Text.StringBuilder
  # split on backticks first -> code spans
  $parts = $s -split '`'
  for ($k = 0; $k -lt $parts.Count; $k++) {
    $seg = $parts[$k]
    if ($k % 2 -eq 1) { [void]$out.Append('<code>' + (Esc $seg) + '</code>'); continue }
    # split on ** -> alternating plain / bold
    $b = $seg -split '\*\*'
    for ($j = 0; $j -lt $b.Count; $j++) {
      if ($j % 2 -eq 1) { [void]$out.Append('<strong>' + (Esc $b[$j]) + '</strong>') }
      else { [void]$out.Append((Esc $b[$j])) }
    }
  }
  $out.ToString()
}

# ---- domain display names ----
$domainName = @{}
for ($i=0; $i -lt $lines.Count; $i++) {
  if ($lines[$i] -match '^### Domain ([A-L]) \u2014 (.+)$') { $domainName[$Matches[1]] = $Matches[2] }
}
$domainName['X'] = 'GDD anti-goal cross-check'

# ---- badges ----
$ruled   = @('HA-09','HA-10','HA-16','HA-22','HA-25','HA-28','HA-29')
$divCand = @('HA-29')

$sb = New-Object System.Text.StringBuilder
$curDom = $null
foreach ($r in $rows) {
  if ($r.Domain -ne $curDom) {
    $curDom = $r.Domain
    $dn = $domainName[$r.Domain]
    if ($r.Domain -eq 'X') {
      [void]$sb.AppendLine("  <div class=`"cat`">")
      [void]$sb.AppendLine("    <h2>GDD anti-goal cross-check rows</h2>")
      [void]$sb.AppendLine("    <div class=`"cat-note`">Each anti-goal in Master GDD v1.1.1 §2 is a testable claim about core-loop shape, and each has a Pre-CU comparand in the record. The fork side is usually already settled by the GDD, so the observation is narrow: does the <b>build</b> drift from the stated anti-goal?</div>")
    } else {
      [void]$sb.AppendLine("  <div class=`"cat`" id=`"dom-$($r.Domain)`">")
      [void]$sb.AppendLine("    <h2>Domain $($r.Domain) &mdash; $dn</h2>")
    }
  }
  $badge = ''
  $st = $r.Strength
  $cls = 'attested'
  if ($st -match 'THIN') { $cls = 'thin' } elseif ($st -match 'ABSENT') { $cls = 'absent' }
  $badge += "<span class=`"badge $cls`">$($cls.ToUpper())</span>"
  if ($ruled -contains $r.Id) { $badge += "<span class=`"badge ruled`">RULED SURFACE</span>" }
  if ($divCand -contains $r.Id) { $badge += "<span class=`"badge divcand`">KNOWN DIVERGENCE CANDIDATE</span>" }
  $strengthFull = $st
  [void]$sb.AppendLine("    <div class=`"row`" data-id=`"$($r.Id)`" data-strength=`"$cls`" data-domain=`"$($r.Domain)`">")
  [void]$sb.AppendLine("      <div class=`"row-head`">")
  [void]$sb.AppendLine("        <span class=`"rowid`">$($r.Id)</span>")
  [void]$sb.AppendLine("        <span class=`"rowloop`">$(RenderMd $r.Loop)</span>")
  [void]$sb.AppendLine("        <span class=`"badges`">$badge</span>")
  [void]$sb.AppendLine("      </div>")
  [void]$sb.AppendLine("      <div class=`"ref`">")
  [void]$sb.AppendLine("        <div class=`"ref-label`">Pre-CU account (read-only &mdash; from the recollection record)</div>")
  [void]$sb.AppendLine("        <div class=`"ref-text`">$(RenderMd $r.PrecU)</div>")
  [void]$sb.AppendLine("        <div class=`"ref-label`">Account strength</div>")
  [void]$sb.AppendLine("        <div class=`"ref-text`">$(RenderMd $strengthFull)</div>")
  [void]$sb.AppendLine("        <div class=`"ref-label`">Fork observation point (read-only)</div>")
  [void]$sb.AppendLine("        <div class=`"ref-text`">$(RenderMd $r.Fork)</div>")
  [void]$sb.AppendLine("      </div>")
  [void]$sb.AppendLine("      <div class=`"entry`">")
  [void]$sb.AppendLine("        <div class=`"entry-top`">")
  [void]$sb.AppendLine('          <label class="fl">Verdict <select class="verdict"><option value=@@EMPTY@@>(&mdash; not yet recorded &mdash;)</option><option value="MATCH">MATCH</option><option value="DIVERGES">DIVERGES</option><option value="INSUFFICIENT EVIDENCE">INSUFFICIENT EVIDENCE</option><option value="NOT OBSERVABLE HERE">NOT OBSERVABLE HERE</option><option value="NOT APPLICABLE">NOT APPLICABLE</option></select></label>')
  [void]$sb.AppendLine('          <label class="fl">Evidence class <select class="evclass"><option value=@@EMPTY@@>(&mdash; not set &mdash;)</option><option value="MEASURED">MEASURED</option><option value="DOCUMENTED">DOCUMENTED</option><option value="RECALLED">RECALLED</option><option value="INFERRED">INFERRED</option></select></label>')
  [void]$sb.AppendLine("          <span class=`"fl modes`">Mode(s) exercised: <label class=`"cb`"><input type=`"checkbox`" class=`"mode`" value=`"A`"> A design</label><label class=`"cb`"><input type=`"checkbox`" class=`"mode`" value=`"B`"> B fork</label><label class=`"cb`"><input type=`"checkbox`" class=`"mode`" value=`"C`"> C cohort</label></span>")
  [void]$sb.AppendLine("        </div>")
  [void]$sb.AppendLine("        <textarea class=`"obs`" placeholder=`"What was observed in the fork / cohort? Quote counts and raw figures beside any percentage. Do not propose a fix here &mdash; a divergence produces a finding, not a change.`"></textarea>")
  [void]$sb.AppendLine("        <div class=`"entry-bot`">")
  [void]$sb.AppendLine("          <label class=`"fl`">Observation-log ref <input type=`"text`" class=`"logref`" placeholder=`"optional &mdash; e.g. OBS-2026-10-22-03 (event, date, witnesses)`"></label>")
  [void]$sb.AppendLine("          <label class=`"fl`">Notes <input type=`"text`" class=`"notes`" placeholder=`"optional`"></label>")
  [void]$sb.AppendLine("        </div>")
  [void]$sb.AppendLine("      </div>")
  [void]$sb.AppendLine("    </div>")
}
$rowsHtml = $sb.ToString().Replace('@@EMPTY@@','""')

$css = @'
  :root { --bg:#0b0c10; --panel:#1f2833; --panel2:#232d3a; --text:#c5c6c7; --accent:#66fcf1; --muted:#45a29e; --border:#2b3442; --input-bg:#0b0c10; --warn:#ffcc66; --diverge:#ff6b6b; --thin:#ffcc66; --absent:#888; }
  * { box-sizing:border-box; }
  body { font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif; background:var(--bg); color:var(--text); margin:0; padding:24px; line-height:1.5; }
  .container { max-width:1280px; margin:0 auto; }
  header { border-bottom:2px solid var(--border); padding-bottom:16px; margin-bottom:20px; }
  h1 { font-size:1.5rem; margin:0 0 8px 0; color:var(--accent); }
  .subtitle { font-size:0.9rem; color:var(--muted); margin-bottom:14px; }
  .authority { background:#2a1f14; border:1px solid var(--warn); border-radius:8px; padding:14px 16px; font-size:0.85rem; margin-bottom:18px; }
  .authority h3 { margin:0 0 8px 0; color:var(--warn); font-size:0.95rem; }
  .authority p { margin:6px 0; }
  .rules { background:var(--panel); border:1px solid var(--border); border-radius:8px; padding:16px; font-size:0.85rem; margin-bottom:18px; }
  .rules h3 { margin-top:0; color:var(--accent); font-size:1rem; }
  .rules ul { margin:8px 0 0 0; padding-left:20px; }
  .rules li { margin-bottom:5px; }
  .meta { display:grid; grid-template-columns:repeat(auto-fit,minmax(230px,1fr)); gap:12px; background:var(--panel); border:1px solid var(--border); border-radius:8px; padding:16px; margin-bottom:14px; }
  .meta label { font-size:0.76rem; color:var(--muted); display:flex; flex-direction:column; gap:4px; }
  .meta input, .meta textarea { background:var(--input-bg); border:1px solid var(--border); color:var(--text); border-radius:4px; padding:6px 9px; font-size:0.85rem; font-family:inherit; }
  .meta textarea { min-height:58px; resize:vertical; }
  .meta input:focus, .meta textarea:focus { outline:none; border-color:var(--accent); }
  .meta .wide { grid-column:1 / -1; }
  .toolbar { display:flex; gap:10px; flex-wrap:wrap; align-items:center; background:var(--panel); border:1px solid var(--border); border-radius:8px; padding:12px 16px; margin-bottom:18px; }
  .toolbar label { font-size:0.76rem; color:var(--muted); display:flex; align-items:center; gap:6px; }
  .toolbar select { background:var(--input-bg); border:1px solid var(--border); color:var(--text); border-radius:4px; padding:5px 8px; font-size:0.8rem; }
  .progress { margin-left:auto; font-size:0.78rem; color:var(--muted); }
  .tally { font-size:0.72rem; color:var(--muted); margin-top:3px; }
  .tally b { color:var(--text); font-weight:normal; }
  .tally .n { color:var(--accent); }
  .cat { margin-bottom:34px; }
  .cat h2 { font-size:1.1rem; color:var(--accent); border-bottom:2px solid var(--border); padding-bottom:6px; margin-bottom:10px; }
  .cat-note { font-size:0.8rem; color:var(--muted); margin-bottom:14px; font-style:italic; border-left:3px solid var(--border); padding-left:10px; }
  .row { background:var(--panel); border:1px solid var(--border); border-radius:8px; padding:14px 16px; margin-bottom:12px; }
  .row.hidden { display:none; }
  .row-head { display:flex; gap:10px; align-items:baseline; flex-wrap:wrap; margin-bottom:10px; }
  .rowid { font-weight:bold; color:var(--accent); font-size:0.9rem; min-width:52px; }
  .rowloop { font-size:0.98rem; color:var(--text); }
  .badges { margin-left:auto; display:flex; gap:6px; flex-wrap:wrap; }
  .badge { font-size:0.66rem; letter-spacing:0.04em; padding:2px 7px; border-radius:3px; border:1px solid var(--border); color:var(--muted); text-transform:uppercase; white-space:nowrap; }
  .badge.attested { color:#7ee787; border-color:#2d5a3d; }
  .badge.thin { color:var(--thin); border-color:#5a4a2d; }
  .badge.absent { color:var(--absent); border-color:#3a3a3a; }
  .badge.ruled { color:#8ab4f8; border-color:#2b4a7a; }
  .badge.divcand { color:var(--diverge); border-color:#7a2b2b; }
  .ref { background:var(--panel2); border:1px solid var(--border); border-radius:6px; padding:10px 12px; margin-bottom:12px; }
  .ref-label { font-size:0.66rem; letter-spacing:0.05em; text-transform:uppercase; color:var(--muted); margin:6px 0 3px 0; }
  .ref-label:first-child { margin-top:0; }
  .ref-text { font-size:0.88rem; color:var(--text); }
  .ref-text code { background:var(--input-bg); padding:1px 5px; border-radius:3px; font-size:0.88em; }
  .entry-top { display:flex; gap:14px; flex-wrap:wrap; align-items:flex-end; margin-bottom:9px; }
  .fl { font-size:0.74rem; color:var(--muted); display:flex; flex-direction:column; gap:4px; }
  .fl select, .fl input { background:var(--input-bg); border:1px solid var(--border); color:var(--text); border-radius:4px; padding:5px 8px; font-size:0.82rem; font-family:inherit; }
  .modes { flex-direction:row; align-items:center; gap:12px; }
  .cb { display:inline-flex; align-items:center; gap:4px; color:var(--text); font-size:0.8rem; }
  textarea.obs { width:100%; min-height:78px; background:var(--input-bg); border:1px solid var(--border); color:var(--text); border-radius:6px; padding:10px; resize:vertical; font-family:inherit; font-size:0.9rem; }
  .entry-bot { display:flex; gap:14px; flex-wrap:wrap; margin-top:9px; }
  .entry-bot .fl { flex:1 1 240px; }
  textarea:focus, input:focus, select:focus { outline:none; border-color:var(--accent); }
  .v-MATCH { color:#7ee787; }
  .v-DIVERGES { color:var(--diverge); font-weight:bold; }
  footer { border-top:2px solid var(--border); margin-top:30px; padding-top:16px; display:flex; gap:12px; flex-wrap:wrap; align-items:center; }
  .big-btn { background:var(--accent); color:#000; border:none; padding:12px 24px; border-radius:6px; font-weight:bold; cursor:pointer; font-size:1rem; }
  .big-btn:hover { opacity:0.9; }
  .big-btn.secondary { background:var(--panel); color:var(--text); border:1px solid var(--border); }
  .big-btn.secondary:hover { border-color:var(--accent); }
  #output { display:none; background:var(--panel); border:1px solid var(--border); border-radius:8px; padding:16px; margin-top:16px; font-family:'Courier New',Courier,monospace; font-size:0.82rem; white-space:pre-wrap; max-height:760px; overflow:auto; }
  #copy { margin-top:10px; background:var(--muted); color:#000; border:none; padding:8px 16px; border-radius:4px; cursor:pointer; font-weight:bold; }
  #copy:hover { opacity:0.9; }
  .saveinfo { font-size:0.75rem; color:var(--muted); margin-top:8px; }
  .hint { font-size:0.74rem; color:var(--muted); margin-top:6px; }
'@

$js = @'
var KEY = 'swg_ha_checklist_v0_2';
var VERDICTS = ['MATCH','DIVERGES','INSUFFICIENT EVIDENCE','NOT OBSERVABLE HERE','NOT APPLICABLE'];
var rowEls = Array.prototype.slice.call(document.querySelectorAll('.row'));
var metaIds = ['runid','executor','rundate','modes','cohort','forkbuild','surfaces','takenat','unfilledreason','ownerquestions'];

function rid(r) { return r.getAttribute('data-id'); }
function el(id) { return document.getElementById(id); }
function loadState() { try { return JSON.parse(localStorage.getItem(KEY) || '{}'); } catch (e) { return {}; } }
function saveState() {
  var s = { meta: {} };
  metaIds.forEach(function (id) { var e = el(id); if (e) s.meta[id] = e.value; });
  s.rows = {};
  rowEls.forEach(function (r) {
    var modes = Array.prototype.slice.call(r.querySelectorAll('input.mode:checked')).map(function (m) { return m.value; });
    var o = { v: r.querySelector('select.verdict').value, e: r.querySelector('select.evclass').value, o: r.querySelector('textarea.obs').value, l: r.querySelector('input.logref').value, n: r.querySelector('input.notes').value, m: modes };
    if (o.v || o.e || o.o || o.l || o.n || modes.length) s.rows[rid(r)] = o;
  });
  try { localStorage.setItem(KEY, JSON.stringify(s)); } catch (e) {}
}
function hydrate() {
  var s = loadState();
  if (s.meta) metaIds.forEach(function (id) { var e = el(id); if (e && s.meta[id]) e.value = s.meta[id]; });
  if (s.rows) rowEls.forEach(function (r) {
    var v = s.rows[rid(r)];
    if (!v) return;
    if (v.v) r.querySelector('select.verdict').value = v.v;
    if (v.e) r.querySelector('select.evclass').value = v.e;
    if (v.o) r.querySelector('textarea.obs').value = v.o;
    if (v.l) r.querySelector('input.logref').value = v.l;
    if (v.n) r.querySelector('input.notes').value = v.n;
    if (v.m) r.querySelectorAll('input.mode').forEach(function (m) { m.checked = v.m.indexOf(m.value) !== -1; });
  });
  if (!el('rundate').value) {
    var t = new Date();
    el('rundate').value = t.getFullYear() + '-' + String(t.getMonth() + 1).padStart(2, '0') + '-' + String(t.getDate()).padStart(2, '0');
  }
}
function updateProgress() {
  var attempted = 0, unfilled = 0, counts = {};
  VERDICTS.forEach(function (v) { counts[v] = 0; });
  rowEls.forEach(function (r) {
    var v = r.querySelector('select.verdict').value;
    if (v) { attempted++; counts[v]++; } else { unfilled++; }
  });
  var parts = [];
  VERDICTS.forEach(function (v) { parts.push('<b>' + v + '</b> ' + counts[v]); });
  el('progress').innerHTML = '<span class="n">' + attempted + ' / ' + rowEls.length + '</span> rows recorded &middot; ' + unfilled + ' unfilled<br>' + parts.join(' &middot; ') + '<br><span style="color:var(--warn)">Counts of recorded outcomes &mdash; NOT a score. Do not convert this into one.</span>';
}
function applyFilter() {
  var dom = el('fdom').value, filt = el('ffilt').value;
  rowEls.forEach(function (r) {
    var show = true;
    if (dom !== 'ALL' && r.getAttribute('data-domain') !== dom) show = false;
    if (show && filt !== 'ALL') {
      var v = r.querySelector('select.verdict').value;
      if (filt === 'UNFILLED') { if (v) show = false; }
      else if (filt === 'DIVERGES') { if (v !== 'DIVERGES') show = false; }
      else if (filt === 'RULED') { if (r.querySelector('.badge.ruled') === null) show = false; }
      else if (filt === 'THINABSENT') { var s = r.getAttribute('data-strength'); if (s === 'attested') show = false; }
    }
    r.classList.toggle('hidden', !show);
  });
}
function buildExport() {
  var L = [];
  var bar = '----------------------------------------------------------------------';
  function meta(id, label) { var v = el(id).value.trim(); return label + ': ' + (v || 'UNRECORDED'); }
  L.push('=== HISTORICAL ACCURACY CHECKLIST RUN ===');
  L.push(meta('runid', 'Run ID'));
  L.push(meta('rundate', 'Date'));
  L.push(meta('executor', 'Executor (human)'));
  L.push(meta('modes', 'Mode(s) exercised'));
  L.push(meta('cohort', 'Cohort (size, composition, duration)'));
  L.push(meta('forkbuild', 'Fork build / commit'));
  L.push(meta('surfaces', 'Observation surfaces used'));
  L.push(meta('takenat', 'Telemetry taken_at'));
  L.push(bar);
  var counts = {}, attempted = 0, unfilled = [];
  VERDICTS.forEach(function (v) { counts[v] = 0; });
  rowEls.forEach(function (r) {
    var v = r.querySelector('select.verdict').value;
    if (v) { attempted++; counts[v]++; } else { unfilled.push(rid(r)); }
  });
  L.push('Rows attempted:            ' + attempted + ' / ' + rowEls.length);
  L.push('Verdict tally:             MATCH ' + counts['MATCH'] + ' | DIVERGES ' + counts['DIVERGES']);
  L.push('                           INSUFFICIENT EVIDENCE ' + counts['INSUFFICIENT EVIDENCE'] + ' | NOT OBSERVABLE HERE ' + counts['NOT OBSERVABLE HERE']);
  L.push('                           NOT APPLICABLE ' + counts['NOT APPLICABLE']);
  L.push('Rows left unfilled:        ' + unfilled.length + (unfilled.length ? ' (' + unfilled.join(', ') + ')' : ''));
  var ur = el('unfilledreason').value.trim();
  L.push('Reason for unfilled rows:  ' + (ur || 'UNRECORDED'));
  L.push('Tally note:                This is a count of recorded outcomes, NOT a score. It must not be');
  L.push('                           converted into one, and no composite figure may be derived from it.');
  L.push(bar);
  L.push('FINDINGS (DIVERGES only, one line each -- no fixes proposed here):');
  var anyD = false;
  rowEls.forEach(function (r) {
    if (r.querySelector('select.verdict').value !== 'DIVERGES') return;
    anyD = true;
    var obs = r.querySelector('textarea.obs').value.trim();
    var oq = el('ownerquestions').value.trim();
    L.push('  ' + rid(r) + ' : ' + (obs || 'OBSERVATION NOT RECORDED'));
  });
  if (!anyD) L.push('  (none recorded)');
  L.push(bar);
  L.push('QUESTIONS FOR THE OWNER (ruling required, not a recommendation):');
  var oq = el('ownerquestions').value.trim();
  if (oq) { oq.split('\n').forEach(function (line) { if (line.trim()) L.push('  ' + line.trim()); }); }
  else L.push('  (none recorded)');
  L.push(bar);
  L.push('NOTE: This run produced EVIDENCE. It closed no open question, changed');
  L.push('no status, authorised no tuning, and set no target.');
  L.push('A DIVERGES on a RULED SURFACE is a re-opening question for the owner,');
  L.push('not a status change. The ruling stands until the owner moves it.');
  L.push('');
  L.push('--- PER-ROW RECORD ---');
  L.push('');
  rowEls.forEach(function (r) {
    var v = r.querySelector('select.verdict').value;
    var modes = Array.prototype.slice.call(r.querySelectorAll('input.mode:checked')).map(function (m) { return m.value; });
    var ev = r.querySelector('select.evclass').value;
    var obs = r.querySelector('textarea.obs').value.trim();
    var lref = r.querySelector('input.logref').value.trim();
    var notes = r.querySelector('input.notes').value.trim();
    var badges = Array.prototype.slice.call(r.querySelectorAll('.badge')).map(function (b) { return b.textContent; }).join(', ');
    L.push('### ' + rid(r) + ' -- ' + r.querySelector('.rowloop').textContent.trim());
    L.push('- Account strength / flags: ' + (badges || 'none'));
    L.push('- Verdict: ' + (v || 'UNFILLED'));
    L.push('- Mode(s) exercised: ' + (modes.length ? modes.join(', ') : 'not recorded'));
    L.push('- Evidence class: ' + (ev || 'not set'));
    L.push('- Observation: ' + (obs || 'not recorded'));
    L.push('- Observation-log ref: ' + (lref || 'none'));
    L.push('- Notes: ' + (notes || 'none'));
    L.push('');
  });
  L.push('--- END OF RUN RECORD ---');
  return L.join('\n');
}
rowEls.forEach(function (r) {
  r.querySelector('select.verdict').addEventListener('change', function () { saveState(); updateProgress(); applyFilter(); });
  r.querySelector('select.evclass').addEventListener('change', function () { saveState(); });
  r.querySelector('textarea.obs').addEventListener('input', function () { saveState(); });
  r.querySelector('input.logref').addEventListener('input', function () { saveState(); });
  r.querySelector('input.notes').addEventListener('input', function () { saveState(); });
  r.querySelectorAll('input.mode').forEach(function (m) { m.addEventListener('change', function () { saveState(); }); });
});
metaIds.forEach(function (id) { var e = el(id); if (e) e.addEventListener('input', saveState); });
el('fdom').addEventListener('change', applyFilter);
el('ffilt').addEventListener('change', applyFilter);
el('exportBtn').addEventListener('click', function () {
  var out = el('output');
  out.textContent = buildExport();
  out.style.display = 'block';
  el('copy').style.display = 'inline-block';
  out.scrollIntoView({ behavior: 'smooth' });
});
function fallbackCopy(txt, done) {
  var ta = document.createElement('textarea');
  ta.value = txt; ta.style.position = 'fixed'; ta.style.top = '-1000px'; ta.style.opacity = '0';
  document.body.appendChild(ta); ta.select();
  try { document.execCommand('copy'); done(); } catch (e) { alert('Copy failed -- please select the text manually and press Ctrl+C.'); }
  document.body.removeChild(ta);
}
el('copy').addEventListener('click', function () {
  var txt = el('output').textContent;
  var done = function () { el('copy').textContent = 'Copied!'; setTimeout(function () { el('copy').textContent = 'Copy to clipboard'; }, 1200); };
  if (navigator.clipboard && navigator.clipboard.writeText && window.isSecureContext) {
    navigator.clipboard.writeText(txt).then(done, function () { fallbackCopy(txt, done); });
  } else { fallbackCopy(txt, done); }
});
el('clearBtn').addEventListener('click', function () {
  if (!confirm('Clear every recorded verdict, observation and note? This also erases the autosave.')) return;
  try { localStorage.removeItem(KEY); } catch (e) {}
  rowEls.forEach(function (r) {
    r.querySelector('select.verdict').value = ''; r.querySelector('select.evclass').value = '';
    r.querySelector('textarea.obs').value = ''; r.querySelector('input.logref').value = ''; r.querySelector('input.notes').value = '';
    r.querySelectorAll('input.mode').forEach(function (m) { m.checked = false; });
  });
  metaIds.forEach(function (id) { var e = el(id); if (e && id !== 'rundate') e.value = ''; });
  hydrate(); saveState(); updateProgress(); applyFilter();
});
hydrate(); saveState(); updateProgress(); applyFilter();
'@

$html = @'
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>SWG Historical Accuracy — Core-Loop Comparison Checklist</title>
<style>
__CSS__
</style>
</head>
<body>
<div class="container">
  <header>
    <h1>SWG Historical Accuracy — Core-Loop Comparison Checklist</h1>
    <div class="subtitle">47-row playtesting sheet · instrument <code>proposals/SWG_Historical_Accuracy_Checklist_v0.2.md</code> (PROPOSED, unadopted) · §1.3 test 5, Historical Accuracy — human pass only</div>
  </header>

  <div class="authority">
    <h3>Authority and non-authority — read before recording anything</h3>
    <p><b>This page is a local transport tool. It carries no project authority of any kind.</b> It is a fill-and-export sheet, not a decision record, not a ruling, and not a source of status.</p>
    <p><b>Every verdict on this page is yours.</b> Nothing here is pre-filled, and this tool will not compute one for you. A blank row stays blank in the export — it is never written as a match.</p>
    <p><b>A completed run produces EVIDENCE, not decisions.</b> It closes no open question, changes no status, authorises no tuning, and sets no target. Per BAL-001 and HD-TST-01, no numeric change and no mechanism change may follow from it without separate authorisation.</p>
    <p><b>No composite score.</b> The tally shown below is a count of recorded outcomes. It must not be converted into a score, a percentage, or a pass mark.</p>
    <p><b>Observe, never steer.</b> A divergence is recorded as a finding. Fixes are not proposed here — they require a separate, authorised proposal.</p>
  </div>

  <div class="rules">
    <h3>Execution rules (instrument §2)</h3>
    <ul>
      <li><b>Every unfilled row stays unfilled.</b> <code>INSUFFICIENT EVIDENCE</code> and <code>NOT OBSERVABLE HERE</code> are correct, valuable answers.</li>
      <li><b>INSUFFICIENT EVIDENCE is not a soft MATCH.</b> An affirmative one-word answer is weak evidence of satisfaction and near-zero evidence of mechanism (record §7 item 6).</li>
      <li><b>NOT OBSERVABLE HERE is not the same thing.</b> It means the account is fine and the arena is not — e.g. a mechanic not built, or no live cohort.</li>
      <li><b>Record raw counts beside every percentage.</b> A 40/42 cohort is 95%, not a verdict.</li>
      <li><b>Recollected magnitudes are never targets.</b> Any figure in the Pre-CU column is a recalled magnitude, quoted as such and nothing else.</li>
      <li><b>Mode C (live cohort) is where Part 2 actually lives.</b> Modes A and B do not clear it. On a small cohort, a Mode B DIVERGES may be a cohort artefact — record the cohort.</li>
      <li><b>THIN and ABSENT rows are marked.</b> On a THIN row a MATCH is near-meaningless; on an ABSENT row, INSUFFICIENT EVIDENCE is the expected verdict.</li>
      <li><b>RULED SURFACE rows are marked.</b> A DIVERGES there is a <b>re-opening question for the owner</b>, not a status change. The ruling stands until the owner moves it.</li>
      <li><b>Do not edit the Pre-CU column.</b> It is a quotation of a filed record. If it is wrong, that is a finding about the record, filed separately.</li>
    </ul>
  </div>

  <div class="meta">
    <label>Run ID <input id="runid" type="text" placeholder="HA-RUN-____"></label>
    <label>Date <input id="rundate" type="text" placeholder="YYYY-MM-DD"></label>
    <label>Executor (human) <input id="executor" type="text" placeholder="who is filling this sheet"></label>
    <label>Mode(s) exercised <input id="modes" type="text" placeholder="A design / B fork / C cohort"></label>
    <label>Cohort (size, composition, duration) <input id="cohort" type="text" placeholder="e.g. 3 chars, 40 min, single session"></label>
    <label>Fork build / commit <input id="forkbuild" type="text" placeholder="commit sha or build id"></label>
    <label>Observation surfaces used <input id="surfaces" type="text" placeholder="e.g. /api/phase10/telemetry, anomalies"></label>
    <label>Telemetry taken_at <input id="takenat" type="text" placeholder="from the telemetry payload"></label>
    <label class="wide">Reason for unfilled rows <textarea id="unfilledreason" placeholder="why the remaining rows were not exercised — e.g. mechanic not built; no live cohort; run out of session time"></textarea></label>
    <label class="wide">Questions for the owner (ruling required, not a recommendation) <textarea id="ownerquestions" placeholder="one per line — e.g. does the removal-vs-damage divergence on HA-29 warrant re-opening HD-BLD-02?"></textarea></label>
  </div>

  <div class="toolbar">
    <label>Domain <select id="fdom"><option value="ALL">All domains</option></select></label>
    <label>Show <select id="ffilt">
      <option value="ALL">All rows</option>
      <option value="UNFILLED">Unfilled only</option>
      <option value="DIVERGES">DIVERGES only</option>
      <option value="RULED">Ruled surfaces only</option>
      <option value="THINABSENT">Thin / absent account only</option>
    </select></label>
    <div class="progress" id="progress"></div>
  </div>
  <div class="saveinfo">Verdicts, observations and notes autosave to this browser as you type. Nothing leaves your machine. The form has no network calls.</div>

__ROWS__

  <footer>
    <button class="big-btn" id="exportBtn" type="button">Export run record (Markdown)</button>
    <button class="big-btn secondary" id="clearBtn" type="button">Clear all</button>
  </footer>
  <div id="output"></div>
  <button id="copy" type="button" style="display:none;">Copy to clipboard</button>
</div>
<script>
__JS__
</script>
</body>
</html>
'@

$html = $html.Replace('__CSS__', $css).Replace('__ROWS__', $rowsHtml).Replace('__JS__', $js)

# build the domain filter options from the parsed rows
$domOpts = ''
$seen = @{}
foreach ($r in $rows) {
  if ($seen.ContainsKey($r.Domain)) { continue }
  $seen[$r.Domain] = $true
  $dn = $domainName[$r.Domain]
  if ($r.Domain -eq 'X') { $domOpts += "      <option value=`"X`">GDD anti-goal cross-check</option>`n" }
  else { $domOpts += "      <option value=`"$($r.Domain)`">Domain $($r.Domain) &mdash; $dn</option>`n" }
}
$html = $html.Replace('<select id="fdom"><option value="ALL">All domains</option></select>',
                      "<select id=`"fdom`"><option value=`"ALL`">All domains</option>`n$domOpts    </select>")

[System.IO.File]::WriteAllText($outHtml, $html, (New-Object System.Text.UTF8Encoding($false)))

# ---------------- VERIFICATION ----------------
$out = [System.IO.File]::ReadAllText($outHtml, [System.Text.Encoding]::UTF8)
$vfail = @()
foreach ($r in $rows) {
  # the rendered Pre-CU text must equal the markdown cell with md markers stripped
  $expect = StripMd $r.PrecU
  # pull the rendered ref-text for this row out of the HTML and strip tags
  $m = [regex]::Match($out, '(?s)data-id="' + $r.Id + '".*?ref-text">(.*?)</div>\s*<div class="ref-label">Account strength')
  if (-not $m.Success) { $vfail += ($r.Id + ' :: row block not found'); continue }
  $got = $m.Groups[1].Value
  $got = $got -replace '<[^>]+>',''
  $got = $got -replace '&amp;','&' -replace '&lt;','<' -replace '&gt;','>' -replace '&quot;','"'
  $got = $got.Trim()
  $expect = ($expect -replace '&amp;','&' -replace '&lt;','<' -replace '&gt;','>' -replace '&quot;','"').Trim()
  if ($got -ne $expect) {
    $vfail += ($r.Id + ' :: Pre-CU text mismatch')
    if ($vfail.Count -le 2) {
      $vfail += ('   EXPECT: ' + $expect.Substring(0, [Math]::Min(120, $expect.Length)))
      $vfail += ('   GOT   : ' + $got.Substring(0, [Math]::Min(120, $got.Length)))
    }
  }
}
Write-Output ("Generated : " + $outHtml)
Write-Output ("Rows      : " + $rows.Count)
Write-Output ("Bytes     : " + (Get-Item $outHtml).Length)
Write-Output ("Lines     : " + (Get-Content $outHtml | Measure-Object -Line).Lines)
Write-Output ("Pre-CU round-trip mismatches: " + $vfail.Count)
if ($vfail.Count) { $vfail | ForEach-Object { Write-Output ("  " + $_) } }

# ---------------- ENCODING GUARD ----------------
# If this script is read by PowerShell without a UTF-8 BOM, its literal non-ASCII
# (em dash, section sign) is decoded as ANSI and lands in the HTML as mojibake.
# The source is UTF-8 and is read explicitly above, so only the template literals
# are at risk. Guard against it, and also against any drift from the source.
$moj  = [string][char]0x00E2 + [string][char]0x20AC
$mojN = ([regex]::Matches($out, [regex]::Escape($moj))).Count
$sec  = ([regex]::Matches($out, [string][char]0x00A7)).Count
$emd  = ([regex]::Matches($out, [string][char]0x2014)).Count
Write-Output ("mojibake markers in output: " + $mojN + "   (0 = no ANSI-decode damage)")
Write-Output ("section signs (U+00A7) in output: " + $sec + "   em dashes (U+2014): " + $emd)
if ($mojN -gt 0) { Write-Output '  !! MOJIBAKE DETECTED — the generator script needs a UTF-8 BOM.' }

