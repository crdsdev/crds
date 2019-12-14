/*
Copyright 2019 The CRDS Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package acronym

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	cw = []string{"cacophonous", "cadaverous", "calamity", "callow", "candid", "capitulate", "capricious", "caricature", "cartographer", "castigate", "catharsis", "caustic", "cease", "cede", "chagrin", "charisma", "charlatan", "chastise", "chimerical", "chronic", "circuitous", "circumlocution", "circumspect", "clandestine", "clemency", "clique", "coercion", "cogent", "cognizant", "colloquial", "collusion", "colossal", "commence", "commiserate", "commodious", "compelling", "compensation", "complacent", "compliant", "composure", "compulsory", "concede", "conceited", "concentric", "conciliatory", "concise", "conclave", "concord", "concurrent", "condone", "confine", "conflagration", "conflate", "confluence", "conformity", "confound", "conglomerate", "conjecture", "connotation", "consensus", "conserve", "consolation", "consolidate", "conspicuous", "consternation", "consummate", "contaminate", "contemplate", "contemporaneous", "contrite", "contrived", "controversial", "conundrum", "converse", "convivial", "copious", "cordial", "correlation", "corroborate", "countenance", "coup", "covert", "coveted", "cower", "craven", "credence", "credible", "crestfallen", "criterion", "cryptic", "culminate", "culpable", "cultivate", "cultivated", "cumbersome", "cumulative", "cursory", "curtail", "cyclical", "cynical"}
	rw = []string{"raconteur", "radical", "rambunctious", "ramification", "rampant", "rancor", "rapport", "rarefied", "rationalization", "ravage", "ravenous", "realm", "reap", "rebellious", "rebuke", "rebuttal", "recalcitrant", "recant", "recapitulate", "recidivism", "recipient", "reciprocate", "recluse", "recoil", "recommence", "recompense", "reconcile", "recondite", "reconnaissance", "recrimination", "rectitude", "redoubtable", "redress", "refined", "refulgent", "refurbish", "refutation", "regime", "regress", "reiterate", "rejuvenate", "relapse", "relegate", "relent", "relentless", "relevance", "relevant", "relinquish", "relish", "remediate", "reminiscent", "remorse", "remote", "remunerate", "renegade", "renege", "renounce", "renowned", "renunciation", "repent", "repercussion", "repertoire", "replenish", "reprehensible", "repress", "reprove", "repudiate", "repugnant", "reputable", "resolute", "resonate", "restitution", "restive", "resurgence", "resuscitate", "retaliate", "reticent", "retort", "retract", "retrench", "retribution", "revelation", "revelry", "reverberate", "revere", "revile", "revoke", "revolutionize", "revulsion", "rhetorical", "rigorous", "riveting", "robust", "rousing", "rudimentary", "ruminate", "rural", "ruse", "rustic", "ruthless"}
	dw = []string{"dally", "dapper", "dauntless", "dawdle", "dearth", "debacle", "debilitate", "debunk", "deduce", "defame", "defiance", "defunct", "dejected", "deleterious", "delicacy", "deluge", "demeanor", "demographic", "denounce", "depict", "deplete", "derivation", "descendant", "descry", "desolate", "destitute", "deter", "detrimental", "devout", "dexterity", "diabolical", "diaphanous", "diatribe", "dichotomy", "didactic", "diffident", "dilettante", "dire", "disconcerting", "discord", "discreet", "discrepancy", "disenfranchise", "disfigure", "disgruntled", "disheveled", "disingenuous", "disinter", "disjointed", "dismal", "dismantle", "dismay", "disparage", "disparity", "dispassionate", "dispatch", "dispel", "dispense", "disperse", "displace", "disposable", "disposition", "disseminate", "dissertation", "dissident", "dissimilar", "dissimulate", "dissipate", "dissipated", "dissipation", "dissolute", "dissolution", "dissolve", "dissonance", "dissonant", "dissuade", "distant", "distend", "distill", "distillation", "distort", "distract", "distraught", "distress", "diverse", "divest", "docile", "dogmatic", "doleful", "domestic", "dominant", "dormant", "dreary", "drudgery", "dubious", "duplicity", "durable", "dwell", "dwindle", "dystopian"}
)

// Random returns a string for the CRD acronym
func Random() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("%s %s %s", cw[rand.Intn(len(cw))], rw[rand.Intn(len(rw))], dw[rand.Intn(len(dw))])
}
