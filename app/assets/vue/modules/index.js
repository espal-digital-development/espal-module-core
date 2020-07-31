/**
 * Add modules here for auto complete to work
 *
 * @typedef Modules
 * @type {Object}
 * @property {import('./user').UserModule} userModule
 */

import * as services from '../services';

// first require only the real modules indexes files
const requireContext = require.context('./', true, /\.\/(?!base)(.*)\/\index.js$/);

/**
 * @type {Modules}
 */
const modules = requireContext.keys().reduce((modules, filename) => {
    // TODO :: this is very strict. Module class names need to be UpperCamelCase
    // create the module name and get the module class name
    const moduleName = filename.split('/')[1] + 'Module';
    const moduleClassName = moduleName[0].toUpperCase() + moduleName.substring(1);
    // require the class
    const ModuleClass = requireContext(filename)[moduleClassName];

    // check which services the class wants
    const servicesToInject = ModuleClass.toString()
        .split('\n')[0]
        .split('(')[1]
        .split(')')[0]
        .split(', ')
        .filter(service => service)
        .map(service => services[service]);

    // create and init the class
    modules[moduleName] = new ModuleClass(...servicesToInject);
    modules[moduleName].init();
    return modules;
}, {});

// Add module name here for exporting
// export the instanced modules here
export const {userModule} = modules;
