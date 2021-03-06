// rptasks.js [2019-07-01 BAR8TL]
// Start report of Tasks being performed
`use strict`;
const sprintf  = require('/nodejs/node_modules/sprintf-js').sprintf;
const sqlite3  = require('/nodejs/node_modules/sqlite3');

module.exports.Drpt = class Drpt {
  constructor() {}

  DspTasks(parm, s) {
    s.SetRunSettings(parm, s);
    for (var rpt of s.Rpt) {
      if (rpt.Pr && rpt.Id == s.Konst.REPORT_EDICRQ) {
        this.buildEdirpt(s, rpt).then((msg) => {
          console.log(sprintf('Report %-8s %s...', 'rptedi', msg));
        });
      }
    }
  }

  buildEdirpt(s, rpt) {
    return new Promise((resolve, reject) => {
      const db = new sqlite3.Database(s.Dbort);
      db.serialize(() => {
        db.run('DELETE FROM rptedi;');
        db.run(s.Sqlst.ISRT_RPTEDI);
        db.close();
      });
      resolve('built');
    });  
  }
}
